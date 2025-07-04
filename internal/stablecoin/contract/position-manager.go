package contract

import (
	"context"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/config"
	"github.com/tsisar/starknet-indexer/internal/utils"
	"math/big"
)

func GetCollateralPools(ctx context.Context, blockNumber uint64, positionID *big.Int) string {
	log.Infof("Geting collateral pools for position ID: %d", positionID)

	contractAddress := config.App.IndexerConfig.Contracts["POSITION_MANAGER"].Address
	log.Debugf("Using COLLATERAL_POOL_CONFIG address: %s", contractAddress)

	low, high := bigIntToFeltU256(positionID)

	result, err := callFunction(ctx, blockNumber, contractAddress, "collateral_pools", low, high)
	if err != nil {
		log.Errorf("Failed to call collateral_pools function: %v", err)
		return ""
	}

	log.Debugf("Raw result from collateral_pools: %v", result)

	pools := normalizeHexString(result[0].String())
	log.Debugf("Collateral pools: %s:", pools)

	return pools
}

func GetPositionAddress(ctx context.Context, blockNumber uint64, positionID *big.Int) string {
	log.Infof("Geting position address: %d", positionID)

	contractAddress := config.App.IndexerConfig.Contracts["POSITION_MANAGER"].Address
	log.Debugf("Using POSITION_MANAGER address: %s", contractAddress)

	low, high := bigIntToFeltU256(positionID)

	result, err := callFunction(ctx, blockNumber, contractAddress, "positions", low, high)
	if err != nil {
		log.Errorf("Failed to call positions function: %v", err)
		return ""
	}

	log.Debugf("Raw result from positions: %v", result)

	address := utils.NormalizeStarkNetAddress(result[0].String())
	log.Debugf("Position address: %s", address)

	return address
}
