package stablecoin

import (
	"context"
	"encoding/json"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/subgraph/collateral_pool_config"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/contract"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
)

func HandleLogInitCollateralPoolId(ctx context.Context, client *ent.Client, event collateral_pool_config.LogInitCollateralPoolId, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)
	poolName := event.Key.CollateralPoolId

	// Check if pool already exists
	_, err := client.Pool.Get(ctx, poolId)
	if err == nil {
		log.Infof("Pool with id %s found", poolId)
		return nil
	}

	log.Infof("Creating new pool with id: %s", poolId)

	// Map values
	debtCeiling := utils.DivByRAD(&event.Data.DebtCeiling)
	liquidationRatio := utils.DivByRAYToDecimal(&event.Data.LiquidationRatio)
	stabilityFeeRate := utils.DivByRAY(&event.Data.StabilityFeeRate)
	debtAccumulatedRate := contract.FetchDebtAccumulatedRate(ctx, transaction.BlockNumber, poolId)

	// Create new Pool
	_, err = client.Pool.Create().
		SetID(poolId).
		SetPoolName(poolName).
		SetDebtCeiling(debtCeiling.String()).                         // DivByRAD(&event.Data.DebtCeiling)
		SetLiquidationRatio(utils.ToDecimalString(liquidationRatio)). // DivByRAYToDecimal(&event.Data.LiquidationRatio)
		SetStabilityFeeRate(stabilityFeeRate.String()).               // divByRAY(&event.Data.StabilityFeeRate)
		SetTokenAdapterAddress(event.Data.Adapter).
		SetLockedCollateral(DefaultPrice).
		SetCollateralPrice(DefaultPrice).
		SetCollateralLastPrice(DefaultPrice).
		SetPriceWithSafetyMargin(DefaultPrice).
		SetRawPrice(DefaultPrice).
		SetTotalBorrowed("0").
		SetTvl("0").
		SetTotalAvailable(debtCeiling.String()). // DivByRAD(&event.Data.DebtCeiling)
		SetDebtAccumulatedRate(utils.ToDecimalString(debtAccumulatedRate)).
		Save(ctx)
	if err != nil {
		log.Errorf("Failed to create pool: %v", err)
		return err
	}

	// Update ProtocolStat
	return updateProtocolStat(ctx, client, poolId)
}

func updateProtocolStat(ctx context.Context, client *ent.Client, poolId string) error {
	stat, err := client.ProtocolStat.Get(ctx, SplyceStatsKey)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Infof("ProtocolStat not found, creating new one")
			pools, _ := json.Marshal([]string{poolId})
			_, err := client.ProtocolStat.Create().
				SetID(SplyceStatsKey).
				SetTotalSupply("0").
				SetTvl("0").
				SetPools(string(pools)).
				Save(ctx)
			return err
		} else {
			log.Errorf("Failed to query ProtocolStat: %v", err)
			return err
		}
	}

	// If ProtocolStat exists
	var poolIDs []string
	if err := json.Unmarshal([]byte(stat.Pools), &poolIDs); err != nil {
		log.Errorf("failed to parse ProtocolStat.Pools: %v", err)
		return err
	}

	// Add new poolId if it doesn't already exist
	for _, id := range poolIDs {
		if id == poolId {
			return nil
		}
	}

	poolIDs = append(poolIDs, poolId)
	newPools, _ := json.Marshal(poolIDs)

	_, err = stat.Update().
		SetPools(string(newPools)).
		Save(ctx)
	return err
}

func HandleLogSetDebtCeiling(ctx context.Context, client *ent.Client, event collateral_pool_config.LogSetDebtCeiling, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	pool, err := client.Pool.Get(ctx, poolId)
	if err != nil {
		log.Errorf("Failed to get pool %s: %v", poolId, err)
		return err
	}

	// debtCeiling = event.params._debtCeiling / RAD
	newDebtCeiling := utils.DivByRAD(&event.Data.DebtCeiling)
	// totalAvailable = debtCeiling
	newTotalAvailable := newDebtCeiling

	_, err = pool.Update().
		SetDebtCeiling(newDebtCeiling.String()).
		SetTotalAvailable(newTotalAvailable.String()).
		Save(ctx)
	if err != nil {
		log.Errorf("Failed to update pool %s: %v", poolId, err)
		return err
	}

	log.Debugf("DebtCeiling updated for pool %s", poolId)
	return nil
}

func HandleSetLiquidationRatio(ctx context.Context, client *ent.Client, event collateral_pool_config.LogSetLiquidationRatio, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	pool, err := client.Pool.Get(ctx, poolId)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Warnf("Pool not found: %s", poolId)
			return nil
		}
		log.Errorf("Failed to get pool %s: %v", poolId, err)
		return err
	}

	// liquidationRatio = data / RAY
	newLiquidationRatio := utils.DivByRAYToDecimal(&event.Data.LiquidationRatio)

	_, err = pool.Update().
		SetLiquidationRatio(utils.ToDecimalString(newLiquidationRatio)).
		Save(ctx)
	if err != nil {
		log.Errorf("Failed to update liquidationRatio for pool %s: %v", poolId, err)
		return err
	}

	log.Debugf("LiquidationRatio updated for pool %s", poolId)
	return nil
}

func HandleSetDebtAccumulatedRate(ctx context.Context, client *ent.Client, event collateral_pool_config.LogSetDebtAccumulatedRate, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	pool, err := client.Pool.Get(ctx, poolId)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Warnf("Pool not found: %s", poolId)
			return nil
		}
		log.Errorf("Failed to get pool %s: %v", poolId, err)
		return err
	}

	// debtAccumulatedRate = debtAccumulatedRate / RAY
	newDebtAccumulatedRate := utils.DivByRAYToDecimal(&event.Data.DebtAccumulatedRate)

	_, err = pool.Update().
		SetDebtAccumulatedRate(utils.ToDecimalString(newDebtAccumulatedRate)).
		Save(ctx)
	if err != nil {
		log.Errorf("Failed to update debtAccumulatedRate for pool %s: %v", poolId, err)
		return err
	}

	log.Debugf("DebtAccumulatedRate updated for pool %s", poolId)
	return nil
}
