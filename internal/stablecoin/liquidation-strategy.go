package stablecoin

import (
	"context"
	"fmt"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/subgraph/liquidation_strategy"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"math/big"
)

func PositionLiquidationHandler(ctx context.Context, client *ent.Client, event liquidation_strategy.LogFixedSpreadLiquidate, transaction model.Transaction) error {
	poolId := utils.EncodeStringToHex(event.Key.CollateralPoolId)
	positionID := utils.NormalizeHash(event.Key.PositionAddress)

	position, err := client.Position.Get(ctx, positionID)
	if err != nil {
		log.Warnf("Position not found: %s", positionID)
		return nil
	}

	// Update locked collateral
	collateralLiquidated := utils.DivByWADToDecimal(&event.Data.CollateralAmountToBeLiquidated)
	currentLockedCollateral, _ := new(big.Float).SetString(position.LockedCollateral)
	newLockedCollateral := new(big.Float).Sub(currentLockedCollateral, collateralLiquidated)

	// Update debt share
	debtShareLiquidated := utils.DivByWADToDecimal(&event.Data.ActualDebtShareToBeLiquidated)
	currentDebtShare, _ := new(big.Float).SetString(position.DebtShare)
	newDebtShare := new(big.Float).Sub(currentDebtShare, debtShareLiquidated)

	// Update debt value
	debtValueLiquidated := utils.DivByRADToDecimal(&event.Data.ActualDebtValueToBeLiquidated)
	currentDebtValue, _ := new(big.Float).SetString(position.DebtValue)
	newDebtValue := new(big.Float).Sub(currentDebtValue, debtValueLiquidated)

	// Check position closure conditions
	closed := false
	if (utils.IsZero(newLockedCollateral) && utils.IsZero(newDebtShare)) || utils.IsZero(newLockedCollateral) {
		closed = true
	}

	if closed {
		_, err := position.Update().
			SetLockedCollateral(utils.ToDecimalString(utils.Zero)).
			SetDebtShare(utils.ToDecimalString(utils.Zero)).
			SetDebtValue(utils.ToDecimalString(utils.Zero)).
			SetSafetyBuffer(utils.ToDecimalString(utils.Zero)).
			SetSafetyBufferInPercent(utils.ToDecimalString(utils.Zero)).
			SetPositionStatus("closed").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to close position: %w", err)
		}

		// decrement user position count
		userID := utils.NormalizeHash(position.UserAddress)
		user, err := client.User.Get(ctx, userID)
		if err == nil {
			newCount := new(big.Int)
			newCount.SetString(user.ActivePositionsCount, 10)
			newCount.Sub(newCount, big.NewInt(1))
			if newCount.Sign() < 0 {
				newCount.SetInt64(0)
			}
			_, _ = user.Update().SetActivePositionsCount(newCount.String()).Save(ctx)
		}

	} else if newLockedCollateral.Sign() > 0 && utils.IsZero(newDebtShare) {
		// All debt repaid but collateral remains
		_, err := position.Update().
			SetLockedCollateral(utils.ToDecimalString(newLockedCollateral)).
			SetDebtShare(utils.ToDecimalString(newDebtShare)).
			SetDebtValue(utils.ToDecimalString(newDebtValue)).
			SetPositionStatus("safe").
			SetSafetyBuffer("1").
			SetSafetyBufferInPercent("1").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to update safe position: %w", err)
		}
	} else {
		// Just update numbers if position stays open
		_, err := position.Update().
			SetLockedCollateral(utils.ToDecimalString(newLockedCollateral)).
			SetDebtShare(utils.ToDecimalString(newDebtShare)).
			SetDebtValue(utils.ToDecimalString(newDebtValue)).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to update position: %w", err)
		}
	}

	// Increment liquidation count
	currentLiquidationCount := new(big.Int)
	currentLiquidationCount.SetString(position.LiquidationCount, 10)
	currentLiquidationCount.Add(currentLiquidationCount, big.NewInt(1))

	_, err = position.Update().SetLiquidationCount(currentLiquidationCount.String()).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update liquidation count: %w", err)
	}

	debtShare := utils.DivByWADToDecimal(&event.Data.ActualDebtShareToBeLiquidated)
	collateralAmount := utils.DivByWADToDecimal(&event.Data.CollateralAmountToBeLiquidated)

	// Create activity "liquidation"
	if err := createPositionActivity(ctx, client, "liquidation", poolId, debtShare, collateralAmount, position, transaction); err != nil {
		return fmt.Errorf("failed to create activity: %w", err)
	}

	return nil
}
