package stablecoin

import (
	"context"
	"fmt"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/subgraph/position_manager"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/contract"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"math/big"
)

func NewPositionHandler(ctx context.Context, client *ent.Client, event position_manager.LogNewPosition, transaction model.Transaction) error {
	// Position identifiers
	positionID := event.Key.PositionId
	userAddress := event.Key.User
	walletAddress := event.Key.Owner

	positionAddress := contract.GetPositionAddress(ctx, &positionID)
	collateralPoolId := contract.GetCollateralPools(ctx, &positionID)

	// Load collateralPool
	collateralPool, err := client.Pool.Get(ctx, collateralPoolId)
	if err != nil {
		log.Warnf("Pool not found: %s", collateralPoolId)
		return nil
	}

	// Create Position entity
	position, err := client.Position.Create().
		SetID(positionAddress).
		SetPositionId(positionID.String()).
		SetPositionAddress(positionAddress).
		SetUserAddress(userAddress).
		SetWalletAddress(walletAddress).
		SetCollateralPool(collateralPool.ID).
		SetCollateralPoolName(collateralPool.PoolName).
		SetLockedCollateral("0").
		SetDebtShare("0").
		SetDebtValue("0").
		SetSafetyBuffer("1").
		SetSafetyBufferInPercent("1").
		SetTvl("0").
		SetPositionStatus("safe").
		SetLiquidationCount("0").
		SetBlockNumber(utils.Uint64ToString(transaction.BlockNumber)).
		SetBlockTimestamp(utils.Uint64ToString(transaction.Timestamp)).
		SetTransaction(transaction.Hash).
		SetPool(collateralPool).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create position: %w", err)
	}

	// Handle Position Activity
	positionResult := contract.GetPositionResult(ctx, collateralPoolId, position.PositionAddress)
	debtShare := utils.DivByWADToDecimal(positionResult.DebtShare)
	collateralAmount := utils.DivByWADToDecimal(positionResult.LockedCollateral)

	// Create position activity state := "created"
	if err := createPositionActivity(ctx, client, "created", collateralPoolId, debtShare, collateralAmount, position, transaction); err != nil {
		log.Errorf("failed to create position activity: %v", err)
	}

	// Handle User entity
	user, err := client.User.Get(ctx, userAddress)
	if err != nil {
		// User not found → create new one
		_, err := client.User.Create().
			SetID(userAddress).
			SetAddress(userAddress).
			SetActivePositionsCount("1").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		// Increment active positions count
		count, _ := new(big.Int).SetString(user.ActivePositionsCount, 10)
		count.Add(count, big.NewInt(1))
		_, err := user.Update().SetActivePositionsCount(count.String()).Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
	}

	return nil
}
