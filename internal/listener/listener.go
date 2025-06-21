package listener

import (
	"context"
	"time"

	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/tsisar/extended-log-go/log"
)

func StartBlockPolling(
	ctx context.Context,
	provider *rpc.Provider,
	interval time.Duration,
	lastSeen uint64,
) <-chan uint64 {
	blockCh := make(chan uint64, 100)

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		defer close(blockCh)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				current, err := provider.BlockNumber(ctx)
				if err != nil {
					log.Warnf("polling block number failed: %v", err)
					continue
				}

				if current > lastSeen {
					for b := lastSeen + 1; b <= current; b++ {
						select {
						case blockCh <- b:
						case <-ctx.Done():
							return
						}
					}
					lastSeen = current
				}
			}
		}
	}()

	return blockCh
}
