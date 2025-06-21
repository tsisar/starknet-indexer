package fetcher

import (
	"context"
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/Tsisar/extended-log-go/log"
)

const ChunkSize = 1000 // Default chunk size for fetching events
const StepBlock = 1000 // Default step size for block range

type Block struct {
	Hash         string
	Number       uint64
	Timestamp    uint64
	Transactions []*felt.Felt
}

func FetchAllEvents(
	ctx context.Context,
	provider *rpc.Provider,
	contract *felt.Felt,
	contractName string,
	fromBlock, toBlock uint64,
) ([]rpc.EmittedEvent, error) {
	var allEvents []rpc.EmittedEvent

	if fromBlock == 0 {
		contractDeploymentBlock, err := FindContractDeploymentBlock(ctx, provider, contract, toBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to find contract deployment block: %w", err)
		}
		log.Infof("[%s] Detected contract deployment block = %d", contractName, fromBlock)
		fromBlock = contractDeploymentBlock
	}

	log.Debugf("[%s] Fetching events for contract %s", contractName, contract.String())

	for start := fromBlock; start <= toBlock; start += StepBlock {
		end := start + StepBlock - 1
		if end > toBlock {
			end = toBlock
		}

		log.Debugf("[%s] Fetching events from block %d to %d...", contractName, start, end)

		input := rpc.EventsInput{
			EventFilter: rpc.EventFilter{
				FromBlock: rpc.WithBlockNumber(start),
				ToBlock:   rpc.WithBlockNumber(end),
				Address:   contract,
			},
			ResultPageRequest: rpc.ResultPageRequest{
				ChunkSize: ChunkSize,
			},
		}

		for {
			res, err := provider.Events(ctx, input)
			if err != nil {
				return nil, fmt.Errorf("fetch events (%d-%d): %w", start, end, err)
			}

			// TODO: Add Block time
			allEvents = append(allEvents, res.Events...)

			if res.ContinuationToken == "" {
				break
			}
			input.ResultPageRequest.ContinuationToken = res.ContinuationToken
		}
	}

	log.Infof("[%s] Total fetched events: %d", contractName, len(allEvents))
	return allEvents, nil
}

func FindContractDeploymentBlock(
	ctx context.Context,
	provider *rpc.Provider,
	contract *felt.Felt,
	latestBlock uint64,
) (uint64, error) {
	low := uint64(0)
	high := latestBlock
	var deploymentBlock uint64 = latestBlock

	for low <= high {
		mid := (low + high) / 2

		_, err := provider.ClassHashAt(ctx, rpc.WithBlockNumber(mid), contract)

		if err != nil {
			// Contract does not exist yet at mid → move forward
			low = mid + 1
		} else {
			// Contract exists → move backward to find first block
			deploymentBlock = mid
			if mid == 0 {
				break
			}
			high = mid - 1
		}
	}

	return deploymentBlock, nil
}

func ExtractBlocks(
	ctx context.Context,
	provider *rpc.Provider,
	events []rpc.EmittedEvent,
) (map[string]Block, error) {
	unique := make(map[string]struct{})
	result := make(map[string]Block)

	for _, ev := range events {
		if ev.BlockHash == nil {
			continue
		}
		hashStr := ev.BlockHash.String()
		if _, exists := unique[hashStr]; exists {
			continue
		}
		unique[hashStr] = struct{}{}

		raw, err := provider.BlockWithTxHashes(ctx, rpc.WithBlockHash(ev.BlockHash))
		if err != nil {
			return nil, fmt.Errorf("get block %s: %w", hashStr, err)
		}

		block, ok := raw.(*rpc.BlockTxHashes)
		if !ok {
			continue
		}

		result[hashStr] = Block{
			Number:       block.Number,
			Hash:         hashStr,
			Timestamp:    block.Timestamp,
			Transactions: block.Transactions,
		}
	}

	return result, nil
}

func (b *Block) GetTransactionIndex(txHash string) (int, error) {
	for i, tx := range b.Transactions {
		if tx.String() == txHash {
			return i, nil
		}
	}
	return -1, fmt.Errorf("transaction %s not found in block %d", txHash, b.Number)
}
