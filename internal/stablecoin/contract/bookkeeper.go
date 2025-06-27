package contract

import (
	"context"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/config"
	"math/big"
)

type PositionResult struct {
	DebtShare        *big.Int
	LockedCollateral *big.Int
}

func GetPositionResult(ctx context.Context, poolId string, positionAddress string) PositionResult {
	log.Infof("Getting position result for pool ID: %s, position address: %s", poolId, positionAddress)

	contractAddress := config.App.IndexerConfig.Contracts["BOOKKEEPER"].Address
	log.Debugf("Using BOOKKEEPER address: %s", contractAddress)

	feltPoolId, err := utils.HexToFelt(poolId)
	if err != nil {
		log.Errorf("Failed to convert pool ID to Felt: %v", err)
		return PositionResult{}
	}

	feltPositionAddress, err := utils.HexToFelt(positionAddress)
	if err != nil {
		log.Errorf("Failed to convert position address to Felt: %v", err)
		return PositionResult{}
	}

	result, err := callFunction(ctx, contractAddress, "positions", feltPoolId, feltPositionAddress)
	if err != nil {
		log.Errorf("Failed to call positions function: %v", err)
		return PositionResult{}
	}

	log.Debugf("Raw result from positions call: %v", result)

	if len(result) != 4 {
		log.Errorf("Invalid result received from contract call")
		return PositionResult{}
	}

	position := PositionResult{
		LockedCollateral: feltU256ToBigInt(result[0], result[1]),
		DebtShare:        feltU256ToBigInt(result[2], result[3]),
	}
	log.Infof("Position result - DebtShare: %v, LockedCollateral: %v",
		position.DebtShare, position.LockedCollateral)

	return position
}
