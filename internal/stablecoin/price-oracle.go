package stablecoin

import (
	"context"
	"github.com/tsisar/starknet-indexer/generated/subgraph/price_consumer"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/contract"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"math/big"

	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
)

func PriceUpdateHandler(ctx context.Context, client *ent.Client, event price_consumer.LogSetPrice, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)

	// Load pool by ID
	pool, err := client.Pool.Get(ctx, poolId)
	if err != nil {
		log.Warnf("Pool not found: %s", poolId)
		return nil
	}

	// Parse current collateralPrice and collateralLastPrice
	collateralPrice, _ := new(big.Float).SetString(pool.CollateralPrice)
	collateralLastPrice, _ := new(big.Float).SetString(pool.CollateralLastPrice)

	// If price was not set yet
	if utils.IsZero(collateralPrice) && utils.IsZero(collateralLastPrice) {
		newPrice := utils.DivByWADToDecimal(&event.Data.RawPrice)
		pool, err = pool.Update().
			SetCollateralPrice(utils.ToDecimalString(newPrice)).
			SetCollateralLastPrice(utils.ToDecimalString(newPrice)).
			Save(ctx)
		if err != nil {
			return err
		}
	} else {
		// Shift prices
		newPrice := utils.DivByWADToDecimal(&event.Data.RawPrice)
		pool, err = pool.Update().
			SetCollateralLastPrice(pool.CollateralPrice).
			SetCollateralPrice(utils.ToDecimalString(newPrice)).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	// Update safety margin, raw price and tvl
	priceWithSafetyMargin := utils.DivByRAYToDecimal(&event.Data.PriceWithSafetyMargin)
	rawPrice := utils.DivByWADToDecimal(&event.Data.RawPrice)
	lockedCollateral, _ := new(big.Float).SetString(pool.LockedCollateral)
	tvl := new(big.Float).Mul(lockedCollateral, collateralPrice)

	pool, err = pool.Update().
		SetPriceWithSafetyMargin(utils.ToDecimalString(priceWithSafetyMargin)).
		SetRawPrice(utils.ToDecimalString(rawPrice)).
		SetTvl(utils.ToDecimalString(tvl)).
		Save(ctx)
	if err != nil {
		return err
	}

	debtAccumulatedRate := contract.FetchDebtAccumulatedRate(ctx, poolId)

	// Update safety buffer for all positions
	positions, err := pool.QueryPositions().All(ctx)
	if err != nil {
		return err
	}

	for _, pos := range positions {
		if pos.DebtShare != "0" && pos.LockedCollateral != "0" {

			lockedCollateral, _ := new(big.Float).SetString(pos.LockedCollateral)
			priceWithSafetyMargin, _ := new(big.Float).SetString(pool.PriceWithSafetyMargin)
			debtShare, _ := new(big.Float).SetString(pos.DebtShare)

			// debtValue = debtShare * debtAccumulatedRate
			debtValue := new(big.Float).Mul(debtShare, debtAccumulatedRate)

			collateralValue := new(big.Float).Mul(lockedCollateral, priceWithSafetyMargin)
			safetyBuffer := new(big.Float).Sub(collateralValue, debtValue)
			if safetyBuffer.Sign() < 0 {
				safetyBuffer = utils.Zero
			}

			safetyBufferPercent := utils.Zero
			if priceWithSafetyMargin.Sign() > 0 && lockedCollateral.Sign() > 0 {
				tmp := new(big.Float).Sub(collateralValue, debtValue)
				tmp.Quo(tmp, priceWithSafetyMargin)
				safetyBufferPercent = new(big.Float).Quo(tmp, lockedCollateral)
			}

			tvl := new(big.Float).Mul(lockedCollateral, collateralPrice)

			status := "safe"
			if utils.IsZero(safetyBuffer) {
				status = "unsafe"
			}

			_, err = pos.Update().
				SetDebtValue(utils.ToDecimalString(debtValue)).
				SetSafetyBuffer(utils.ToDecimalString(safetyBuffer)).
				SetSafetyBufferInPercent(utils.ToDecimalString(safetyBufferPercent)).
				SetPositionStatus(status).
				SetTvl(utils.ToDecimalString(tvl)).
				Save(ctx)
			if err != nil {
				log.Errorf("Failed to update position %s: %v", pos.ID, err)
			}
		}
	}

	if err := updateProtocolTVL(ctx, client); err != nil {
		log.Errorf("Failed to update protocol TVL: %v", err)
	}

	return nil
}
