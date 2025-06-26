package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/indexer"
	"github.com/tsisar/starknet-indexer/internal/config"
	"github.com/tsisar/starknet-indexer/internal/fetcher"
	"github.com/tsisar/starknet-indexer/internal/listener"
	"github.com/tsisar/starknet-indexer/internal/storage"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

var status = model.Status{}

func main() {
	g, err := storage.InitGorm()
	if err != nil {
		log.Fatalf("[main] Failed to init Gorm DB: %v", err)
	}
	defer g.Close()

	ctx := context.Background()

	provider, err := rpc.NewProvider(config.App.RPCEndpoint)
	if err != nil {
		log.Fatalf("create provider: %v", err)
	}

	if ok, err := status.Load(ctx, g.DB); err != nil {
		log.Fatalf("failed to load status: %v", err)
	} else if !ok {
		status.Init()
		if err = status.Save(ctx, g.DB); err != nil {
			log.Fatalf("failed to save status: %v", err)
		}
	}

	latestBlockNumber, err := provider.BlockNumber(ctx)
	if err != nil {
		setError(ctx, g.DB, fmt.Sprintf("failed to get latest block number: %v", err))
		log.Fatalf("failed to get latest block number: %v", err)
	}

	fromBlock := config.App.StartBlock
	contracts := config.App.Contracts

	fetchEventsForAllContracts(ctx, g.DB, provider, contracts, fromBlock, latestBlockNumber)

	blockCh := listener.StartBlockPolling(ctx, provider, 10*time.Second, latestBlockNumber)
	setSynced(ctx, g.DB, true)

	for blockNumber := range blockCh {
		fetchEventsForAllContracts(ctx, g.DB, provider, contracts, blockNumber, blockNumber)
		updateCurrentBlock(ctx, g.DB, blockNumber)
	}
}

func fetchEventsForAllContracts(ctx context.Context, db *gorm.DB, provider *rpc.Provider, contracts map[string]string, fromBlock, toBlock uint64) {
	const maxRetries = 3
	const retryDelay = 2 * time.Second

	var wg sync.WaitGroup

	for contractName, addressString := range contracts {
		wg.Add(1)

		go func(name, address string) {
			defer wg.Done()

			var err error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				err = fetchEventsForContract(ctx, db, provider, name, address, fromBlock, toBlock)
				if err == nil {
					return
				}

				log.Warnf("Attempt %d/%d failed for contract %s: %v", attempt, maxRetries, name, err)
				time.Sleep(retryDelay)
			}

			setError(ctx, db, fmt.Sprintf("Failed to fetch events for contract %s at address %s after %d attempts: %v", name, address, maxRetries, err))
			log.Fatalf("All attempts failed for contract %s: %v", name, err)
		}(contractName, addressString)
	}

	wg.Wait()
}

func fetchEventsForContract(ctx context.Context, db *gorm.DB, provider *rpc.Provider, contractName, addressString string, fromBlock, toBlock uint64) error {
	//fmt.Println("---------------------------------")

	cleanName := strings.ToLower(contractName)
	log.Infof("[%s] Fetching events for address %s", cleanName, addressString)

	address, err := utils.HexToFelt(addressString)
	if err != nil {
		return fmt.Errorf("invalid address %s: %w", addressString, err)
	}

	events, err := fetcher.FetchAllEvents(ctx, provider, address, cleanName, fromBlock, toBlock)
	if err != nil {
		return fmt.Errorf("fetch events for contract %s: %w", contractName, err)
	}
	log.Debugf("[%s] Fetched %d events", cleanName, len(events))

	if len(events) == 0 {
		log.Infof("[%s] No events found for contract %s in blocks %d to %d", cleanName, addressString, fromBlock, toBlock)
		return nil
	}

	blocks, err := fetcher.ExtractBlocks(ctx, provider, events)
	if err != nil {
		return fmt.Errorf("extract block timestamps for contract %s: %w", contractName, err)
	}

	decoderMap, ok := generated.Decoders[cleanName]
	if !ok {
		return fmt.Errorf("no decoders found for contract %s", cleanName)
	}

	for index, ev := range events {
		if len(ev.Keys) == 0 {
			log.Warnf("[%s] Skipping event with no keys", cleanName)
			continue
		}

		selector := ev.Keys[0].String()
		decoder, found := decoderMap[selector]
		if !found {
			log.Warnf("[%s] No decoder for event %s in contract %s", cleanName, selector, cleanName)
			continue
		}

		decoded, err := decoder.Func(ev.Keys[1:], ev.Data)
		if err != nil {
			return fmt.Errorf("failed to decode event %s: %w", selector, err)
		}

		log.Debugf("[%s] Blocks: %v", cleanName, blocks)
		block := blocks[ev.BlockHash.String()]
		if err = saveParsedEvent(ctx, db, cleanName, addressString, decoder.FullName(), index, decoded, block, ev.TransactionHash.String()); err != nil {
			return fmt.Errorf("failed to save parsed event %s: %w", decoder.FullName(), err)
		}
		logDecodedEvent(cleanName, addressString, decoder.FullName(), index, decoded, block, ev.TransactionHash.String())
		updateHash(ctx, db, block.Hash, block.Timestamp)
	}
	return nil
}

func logDecodedEvent(
	contractName, contractAddress, eventName string,
	eventIndex int,
	event interface{},
	block fetcher.Block,
	txHash string,
) {
	txIndex, err := block.GetTransactionIndex(txHash)
	if err != nil {
		log.Errorf("Failed to get transaction index for %s in block %s: %v", txHash, block.Hash, err)
		return
	}

	entry := map[string]interface{}{
		"contract_name":    contractName,
		"contract_address": contractAddress,
		"name":             eventName,
		"index":            eventIndex,
		"event":            event,
		"block_number":     block.Number,
		"block_timestamp":  block.Timestamp,
		"block_hash":       block.Hash,
		"tx_index":         txIndex,
		"tx_hash":          txHash,
	}

	jsonBytes, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Errorf("Failed to marshal decoded event %s: %v", eventName, err)
		return
	}

	log.Infof("[%s] Decoded event:\n%s", contractName, string(jsonBytes))
}

func saveParsedEvent(
	ctx context.Context,
	db *gorm.DB,
	contractName, contractAddress, eventName string,
	eventIndex int,
	event interface{},
	block fetcher.Block,
	txHash string,
) error {
	contract := model.Contract{
		Address: contractAddress,
		Name:    contractName,
	}
	if ok, err := contract.Load(ctx, db); err != nil {
		return fmt.Errorf("load contract %s: %w", contractAddress, err)
	} else if !ok {
		if err = contract.Save(ctx, db); err != nil {
			return fmt.Errorf("save contract %s: %w", contractAddress, err)
		}
	}

	txIndex, err := block.GetTransactionIndex(txHash)
	if err != nil {
		return fmt.Errorf("get transaction index for %s in block %s: %w", txHash, block.Hash, err)
	}

	b := model.Block{
		Hash:      block.Hash,
		Timestamp: block.Timestamp,
		Number:    block.Number,
	}
	if ok, err := b.Load(ctx, db); err != nil {
		return fmt.Errorf("load block %s: %w", block.Hash, err)
	} else if !ok {
		if err = b.Save(ctx, db); err != nil {
			return fmt.Errorf("save block %s: %w", block.Hash, err)
		}
	}

	tx := model.Transaction{
		Hash:        txHash,
		Index:       txIndex,
		BlockNumber: block.Number,
		Timestamp:   block.Timestamp,
	}
	if ok, err := tx.Load(ctx, db); err != nil {
		return fmt.Errorf("load transaction %s: %w", txHash, err)
	} else if !ok {
		if err = tx.Save(ctx, db); err != nil {
			return fmt.Errorf("save transaction %s: %w", txHash, err)
		}
	}

	jsonBytes, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal event %s: %w", eventName, err)
	}

	indexStr := fmt.Sprintf("%d", eventIndex)
	jsonStr := string(jsonBytes)

	id := generateId(contractAddress, block.Hash, txHash, indexStr, jsonStr)
	e := model.Event{
		ID:              id,
		Index:           eventIndex,
		Name:            eventName,
		TxIndex:         tx.Index,
		TxHash:          tx.Hash,
		BlockNumber:     block.Number,
		ContractAddress: contract.Address,
		JsonEv:          datatypes.JSON(jsonBytes),
		Timestamp:       block.Timestamp,
	}
	ok, err := e.Load(ctx, db)
	if err != nil {
		return fmt.Errorf("load event %s: %w", eventName, err)
	} else if ok {
		log.Warnf("Event %s already exists", eventName)
	}
	if err = e.Save(ctx, db); err != nil {
		return fmt.Errorf("save event %s: %w", eventName, err)
	}

	return nil
}

// generateId creates a SHA-256 hash from all input strings concatenated in order.
// Returns the result as a hex-encoded string.
func generateId(parts ...string) string {
	hasher := sha256.New()
	for _, s := range parts {
		hasher.Write([]byte(s))
	}
	sum := hasher.Sum(nil)
	return hex.EncodeToString(sum)
}

func updateCurrentBlock(ctx context.Context, db *gorm.DB, blockNumber uint64) {
	if err := status.UpdateCurrentBlock(ctx, db, blockNumber); err != nil {
		log.Errorf("failed to update current block number: %v", err)
	}
}

func updateHash(ctx context.Context, db *gorm.DB, hash string, timestamp uint64) {
	if err := status.UpdateHash(ctx, db, hash, timestamp); err != nil {
		log.Errorf("failed to update current block number: %v", err)
	}
}

func setSynced(ctx context.Context, db *gorm.DB, synced bool) {
	if err := status.SetSynced(ctx, db, synced); err != nil {
		log.Errorf("failed to set synced status: %v", err)
	}
}

func setError(ctx context.Context, db *gorm.DB, msg string) {
	if err := status.SetError(ctx, db, true, msg); err != nil {
		log.Errorf("failed to set error status: %v", err)
	}
}
