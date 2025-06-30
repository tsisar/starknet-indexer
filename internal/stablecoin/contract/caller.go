package contract

import (
	"context"
	"fmt"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/config"
	"math/big"
	"strings"
)

func getRrcProvider() (*rpc.Provider, error) {
	client, err := rpc.NewProvider(config.App.RPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to reconnect to RPC provider: %w", err)
	}

	return client, nil
}

func callFunction(ctx context.Context, blockNumber uint64, contractAddress, funcName string, callData ...*felt.Felt) ([]*felt.Felt, error) {
	client, err := getRrcProvider()
	if err != nil {
		log.Errorf("failed to get RPC provider: %v", err)
	}

	feltContractAddress, err := utils.HexToFelt(contractAddress) // Contract address
	if err != nil {
		log.Errorf("failed to convert contract address to felt: %v", err)
	}

	selector := utils.GetSelectorFromNameFelt(funcName)

	fnCall := rpc.FunctionCall{
		ContractAddress:    feltContractAddress,
		EntryPointSelector: selector,
		Calldata:           callData,
	}

	result, err := client.Call(ctx, fnCall, rpc.BlockID{Number: &blockNumber})
	if err != nil {
		log.Errorf("failed to call contract: %v", err)
	}

	if len(result) == 0 {
		log.Errorf("empty response from contract call")
	}

	return result, nil
}

// feltU256ToBigInt converts StarkNet U256 (low, high) to big.Int
func feltU256ToBigInt(low *felt.Felt, high *felt.Felt) *big.Int {
	lowBI := utils.FeltToBigInt(low)
	highBI := utils.FeltToBigInt(high)

	highBI.Lsh(highBI, 128)
	return highBI.Add(highBI, lowBI)
}

// bigIntToFeltU256 converts big.Int to StarkNet U256 (low, high)
func bigIntToFeltU256(value *big.Int) (*felt.Felt, *felt.Felt) {
	mask := new(big.Int).Lsh(big.NewInt(1), 128)
	mask.Sub(mask, big.NewInt(1))

	low := new(big.Int).And(value, mask)
	high := new(big.Int).Rsh(value, 128)

	lowFelt := utils.BigIntToFelt(low)
	highFelt := utils.BigIntToFelt(high)

	return lowFelt, highFelt
}

func normalizeHexString(input string) string {
	if len(input) < 2 || input[:2] != "0x" {
		return strings.ToUpper(input)
	}
	return "0x" + strings.ToUpper(input[2:])
}
