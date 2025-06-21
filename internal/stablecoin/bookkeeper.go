package stablecoin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tsisar/extended-log-go/log"
	"github.com/Tsisar/starknet-indexer/generated/ent"
	"github.com/Tsisar/starknet-indexer/generated/subgraph/bookkeeper"
	"github.com/Tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/Tsisar/starknet-indexer/internal/storage/model"
	"math/big"
)

func AdjustPositionHandler(ctx context.Context, client *ent.Client, event bookkeeper.LogAdjustPosition, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	// Lookup Pool in the database
	pool, err := client.Pool.Get(ctx, poolId)
	if err == nil {
		if err := updatePool(ctx, pool, event); err != nil {
			log.Errorf("Failed to update pool %s: %v", poolId, err)
		}
	}

	// Update the total TVL in protcol by adding the TVLs from all pools
	if err := updateProtocolTVL(ctx, client); err != nil {
		log.Errorf("Failed to update ProtocolStat TVL: %v", err)
	}

	// Update the position
	if err := updatePosition(ctx, client, event.Key.PositionAddress, event, pool, transaction); err != nil {
		log.Errorf("Failed to update position %s: %v", event.Key.PositionAddress, err)
	}

	return nil
}

func updatePool(ctx context.Context, pool *ent.Pool, position bookkeeper.LogAdjustPosition) error {
	// Convert added collateral: _addCollateral / WAD
	addCollateral := utils.DivByWADToDecimal(&position.Data.AddCollateral)

	// Compute new lockedCollateral = lockedCollateral + addCollateral
	currentLockedCollateral, _ := new(big.Float).SetString(pool.LockedCollateral)
	newLockedCollateral := new(big.Float).Add(currentLockedCollateral, addCollateral)

	// Compute new totalAvailable = debtCeiling - totalBorrowed
	debtCeiling, _ := new(big.Float).SetString(pool.DebtCeiling)
	totalBorrowed, _ := new(big.Float).SetString(pool.TotalBorrowed)
	newTotalAvailable := new(big.Float).Sub(debtCeiling, totalBorrowed)

	// Compute new TVL = lockedCollateral * collateralPrice
	collateralPrice, _ := new(big.Float).SetString(pool.CollateralPrice)
	newTVL := new(big.Float).Mul(newLockedCollateral, collateralPrice)

	// Save updated values
	_, err := pool.Update().
		SetLockedCollateral(utils.ToDecimalString(newLockedCollateral)).
		SetTotalAvailable(utils.ToDecimalString(newTotalAvailable)).
		SetTvl(utils.ToDecimalString(newTVL)).
		Save(ctx)

	if err != nil {
		log.Errorf("Failed to update pool: %v", err)
		return err
	}

	log.Debugf("Pool updated: %s", pool.ID)
	return nil
}

func updateProtocolTVL(ctx context.Context, client *ent.Client) error {
	// Update the total TVL in protcol by adding the TVLs from all pools
	stats, err := client.ProtocolStat.Get(ctx, SplyceStatsKey)
	if err != nil {
		return fmt.Errorf("failed to get ProtocolStat: %w", err)
	}

	var poolIDs []string
	if err := json.Unmarshal([]byte(stats.Pools), &poolIDs); err != nil {
		log.Errorf("Failed to unmarshal pools: %v", err)
		return err
	}

	aggregatedTVL := new(big.Float).Set(utils.Zero)

	for _, poolID := range poolIDs {
		log.Debugf("Processing pool: %s", poolID)

		pool, err := client.Pool.Get(ctx, poolID)
		if err != nil {
			if ent.IsNotFound(err) {
				log.Warnf("Pool not found: %s", poolID)
				continue
			}
			log.Errorf("Failed to get pool %s: %v", poolID, err)
			return err
		}

		tvl, ok := new(big.Float).SetString(pool.Tvl)
		if !ok {
			log.Warnf("Invalid TVL for pool %s: %s", pool.ID, pool.Tvl)
			continue
		}

		aggregatedTVL.Add(aggregatedTVL, tvl)
	}

	_, err = stats.Update().SetTvl(utils.ToDecimalString(aggregatedTVL)).Save(ctx)
	if err != nil {
		log.Errorf("Failed to update ProtocolStat TVL: %v", err)
		return err
	}

	log.Infof("ProtocolStat TVL updated: %s => %s", SplyceStatsKey, utils.ToDecimalString(aggregatedTVL))
	return nil
}

func updatePosition(ctx context.Context, client *ent.Client, positionID string, event bookkeeper.LogAdjustPosition, pool *ent.Pool, transaction model.Transaction) error {
	// Load position
	position, err := client.Position.Get(ctx, positionID)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Warnf("Position not found: %s", positionID)
			return nil
		}
		return fmt.Errorf("failed to load position: %w", err)
	}

	if pool == nil {
		log.Warnf("Pool not found for position %s", positionID)
		return nil
	}

	// lockedCollateral = event.params._lockedCollateral / WAD
	lockedCollateral := utils.DivByWADToDecimal(&event.Data.LockedCollateral)
	debtValue := utils.DivByRADToDecimal(&event.Data.PositionDebtValue)
	debtShare := utils.DivByWADToDecimal(&event.Data.DebtShare)

	// tvl = lockedCollateral * pool.collateralPrice
	collateralPrice, _ := new(big.Float).SetString(pool.CollateralPrice)
	tvl := new(big.Float).Mul(lockedCollateral, collateralPrice)

	// Update fields
	positionUpdate := position.Update().
		SetLockedCollateral(utils.ToDecimalString(lockedCollateral)).
		SetDebtValue(utils.ToDecimalString(debtValue)).
		SetDebtShare(utils.ToDecimalString(debtShare)).
		SetTvl(utils.ToDecimalString(tvl))

	// If debtShare == 0 and position not closed — close it
	if utils.IsZero(debtShare) && position.PositionStatus != "closed" {
		positionUpdate.
			SetPositionStatus("closed").
			SetSafetyBuffer("0").
			SetSafetyBufferInPercent("0").
			SetDebtValue("0")

		// Decrement user's activePositionsCount
		user, err := client.User.Get(ctx, position.UserAddress)
		if err == nil {
			activeCount, _ := new(big.Int).SetString(user.ActivePositionsCount, 10)
			activeCount.Sub(activeCount, big.NewInt(1))
			_, _ = user.Update().SetActivePositionsCount(activeCount.String()).Save(ctx)
		}
	}

	// Update safety buffer
	priceWithSafetyMargin, _ := new(big.Float).SetString(pool.PriceWithSafetyMargin)
	if priceWithSafetyMargin.Cmp(utils.Zero) > 0 && lockedCollateral.Cmp(utils.Zero) > 0 {
		// collateralAvailableToWithdraw = ((priceWithSafetyMargin * lockedCollateral) - debtValue) / priceWithSafetyMargin
		tmp := new(big.Float).Mul(priceWithSafetyMargin, lockedCollateral)
		tmp.Sub(tmp, debtValue)
		collateralAvailableToWithdraw := new(big.Float).Quo(tmp, priceWithSafetyMargin)

		safetyBufferPercent := new(big.Float).Quo(collateralAvailableToWithdraw, lockedCollateral)
		positionUpdate.SetSafetyBufferInPercent(utils.ToDecimalString(safetyBufferPercent))
	}

	// Save updated position
	if _, err := positionUpdate.Save(ctx); err != nil {
		return fmt.Errorf("failed to update position: %w", err)
	}

	log.Infof("Updated position: %s", positionID)

	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	debtShareAdded := utils.DivByWADToDecimal(&event.Data.AddDebtShare)
	collateralAdded := utils.DivByWADToDecimal(&event.Data.AddCollateral)

	//derive the activity state
	state := "topup"
	if utils.IsZero(debtShare) {
		//if debtShareAdded is zero, it means position is closed
		state = "closed"
	} else if debtShareAdded.Cmp(big.NewFloat(0)) < 0 {
		//if debtShareAdded is negative, it means debt is repaid
		state = "repay"
	}

	// TODO: Check if is it correct to skip activity creation if position is closed
	//if state == "closed" {
	//	log.Infof("Skipping activity creation since position is closed: %s", position.ID)
	//	return nil
	//}

	// TODO: Check if debtShareAdded is ok, or need to use debtShare for debtAccumulatedRate
	if err := createPositionActivity(ctx, client, state, poolId, debtShareAdded, collateralAdded, position, transaction); err != nil {
		return fmt.Errorf("failed to create activity: %w", err)
	}

	return nil
}

func SetTotalDebtCeilingHandler(ctx context.Context, client *ent.Client, event bookkeeper.LogSetTotalDebtCeiling, transaction model.Transaction) error {
	stats, err := client.ProtocolStat.Get(ctx, SplyceStatsKey)
	if err != nil {
		// ProtocolStat not found, create a new one
		poolsJSON, _ := json.Marshal([]string{})
		_, err := client.ProtocolStat.Create().
			SetID(SplyceStatsKey).
			SetTvl("0").
			SetTotalSupply("0").
			SetPools(string(poolsJSON)).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create new ProtocolStat: %w", err)
		}
		return nil
	}

	// Update TotalSupply
	totalSupply := utils.DivByRAD(&event.Data.TotalDebtCeiling)
	_, err = stats.Update().
		SetTotalSupply(totalSupply.String()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update TotalSupply: %w", err)
	}

	return nil
}

func StablecoinIssuedAmountHandler(ctx context.Context, client *ent.Client, event bookkeeper.StablecoinIssuedAmount, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	// Load pool
	pool, err := client.Pool.Get(ctx, poolId)
	if err != nil {
		log.Warnf("Pool not found: %s", poolId)
		return nil
	}

	// Convert _poolStablecoinIssued using DivByRAD
	totalBorrowed := utils.DivByRAD(&event.Data.PoolStablecoinIssued)

	// Convert existing debtCeiling and calculate totalAvailable
	debtCeiling, ok := new(big.Int).SetString(pool.DebtCeiling, 10)
	if !ok {
		log.Warnf("Invalid debtCeiling for pool %s: %s", poolId, pool.DebtCeiling)
		return nil
	}
	totalAvailable := new(big.Int).Sub(debtCeiling, totalBorrowed)

	// Update pool fields
	_, err = pool.Update().
		SetTotalBorrowed(totalBorrowed.String()).
		SetTotalAvailable(totalAvailable.String()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed to update pool: %w", err)
	}

	return nil
}
