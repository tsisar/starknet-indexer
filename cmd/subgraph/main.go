package main

import (
	"context"
	"fmt"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/graphql"
	"github.com/tsisar/starknet-indexer/internal/mapper"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/lib/pq"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/storage"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"github.com/tsisar/starknet-indexer/internal/trigger"
	"gorm.io/gorm"
)

var (
	blockQueue   = make(chan uint64, 1000)
	queueRWMutex sync.RWMutex
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, err := storage.InitGorm()
	if err != nil {
		log.Fatalf("[main] Failed to init Gorm DB: %v", err)
	}
	defer g.Close()

	e, err := storage.InitEnt(ctx)
	if err != nil {
		log.Fatalf("[main] Failed to init ent DB: %v", err)
	}
	defer e.Close()

	sqlDB, err := g.DB.DB()
	if err != nil {
		log.Fatalf("[main] Failed to get raw DB: %v", err)
	}

	if err := storage.TruncateAllTables(ctx, e); err != nil {
		log.Fatalf("[main] Failed to truncate all tables: %v", err)
	}

	if err := trigger.SetupCurrentBlockTrigger(sqlDB); err != nil {
		log.Fatalf("setup current block trigger failed: %v", err)
	}
	if err := trigger.SetupSyncedTrigger(sqlDB); err != nil {
		log.Fatalf("setup synced trigger failed: %v", err)
	}

	status := &model.Status{ID: 1}
	if found, err := status.Load(ctx, g.DB); err != nil {
		log.Fatalf("failed to load status: %v", err)
	} else if found && status.Synced && !status.HasError {
		if err := getAllBlocks(ctx, g.DB); err != nil {
			log.Fatalf("getAllBlocks failed: %v", err)
		}
	}

	go processBlockQueue(ctx, g.DB, e)

	go func() {
		if err := listenNotifications(ctx, storage.DSNFromConfig(), handleNotification(ctx, g.DB, e)); err != nil {
			log.Errorf("listen failed: %v", err)
		}
	}()

	if err := graphql.Start(ctx, g.DB, e); err != nil {
		log.Fatalf("GraphQL server failed: %v", err)
	}
}

func handleNotification(ctx context.Context, db *gorm.DB, e *ent.Client) func(channel, payload string) {
	return func(channel, payload string) {
		switch channel {
		case "starknet_current_block_update":
			blockNum, err := strconv.ParseUint(payload, 10, 64)
			if err != nil {
				log.Errorf("Invalid block number: %v", err)
				return
			}

			queueRWMutex.RLock()
			blockQueue <- blockNum
			queueRWMutex.RUnlock()

		case "starknet_synced_update":
			if payload == "true" {
				if err := getAllBlocks(ctx, db); err != nil {
					log.Fatalf("getAllBlocks failed: %v", err)
				}
			}
		}
	}
}

func listenNotifications(ctx context.Context, dsn string, logFn func(channel, payload string)) error {
	listener := pq.NewListener(dsn, 10*time.Second, time.Minute, nil)

	if err := listener.Listen("starknet_current_block_update"); err != nil {
		return fmt.Errorf("listen starknet_current_block_update: %w", err)
	}
	if err := listener.Listen("starknet_synced_update"); err != nil {
		return fmt.Errorf("listen starknet_synced_update: %w", err)
	}

	log.Debug("Listening for PostgreSQL NOTIFY on current_block/synced...")

	for {
		select {
		case <-ctx.Done():
			_ = listener.UnlistenAll()
			return ctx.Err()
		case n := <-listener.Notify:
			if n != nil {
				logFn(n.Channel, n.Extra)
			}
		}
	}
}

func getAllBlocks(ctx context.Context, db *gorm.DB) error {
	queueRWMutex.Lock()
	defer queueRWMutex.Unlock()

	blocks, err := model.GetAllBlocks(ctx, db)
	if err != nil {
		return err
	}

	for _, block := range blocks {
		blockQueue <- block.Number
	}

	return nil
}

func processBlockQueue(ctx context.Context, db *gorm.DB, e *ent.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case blockNumber := <-blockQueue:
			log.Infof("Processing block %d", blockNumber)

			events, err := model.GetEventsByBlock(ctx, db, blockNumber)
			if err != nil {
				log.Errorf("Failed to get events for block %d: %v", blockNumber, err)
				continue
			}

			for _, ev := range events {
				if err := mapper.MapEvents(ctx, e, ev); err != nil {
					log.Errorf("Failed to map event %d in block %d: %v", ev.Index, blockNumber, err)
				}
			}
		}
	}
}
