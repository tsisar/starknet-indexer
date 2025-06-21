package stablecoin

import (
	"context"
	"fmt"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/contract"
	"github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"math/big"
)

// createPositionActivity creates a PositionActivity record
func createPositionActivity(ctx context.Context, client *ent.Client, state, poolId string, debtShare, collateralAmount *big.Float, position *ent.Position, transaction model.Transaction) error {
	activityID := fmt.Sprintf("%s-%s", PositionActivityPrefixKey, utils.NormalizeHash(transaction.Hash))

	debtAccumulatedRate := big.NewFloat(1)
	if !utils.IsZero(debtShare) {
		debtAccumulatedRate = contract.FetchDebtAccumulatedRate(ctx, poolId)
	}
	debtAmount := new(big.Float).Mul(debtShare, debtAccumulatedRate)

	_, err := client.PositionActivity.Get(ctx, activityID)
	if err == nil {
		log.Infof("PositionActivity already exists: %s", activityID)
		return nil
	}

	log.Debugf("PositionActivity not found, creating new one: %s", activityID)

	_, err = client.PositionActivity.Create().
		SetID(activityID).
		SetActivityState(state).
		SetCollateralAmount(utils.ToDecimalString(collateralAmount)).
		SetDebtAmount(utils.ToDecimalString(debtAmount)).
		SetPosition(position).
		SetBlockNumber(utils.Uint64ToString(transaction.BlockNumber)).
		SetBlockTimestamp(utils.Uint64ToString(transaction.Timestamp)).
		SetTransaction(transaction.Hash).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed to create position activity: %w", err)
	}

	log.Debugf("Position activity %s created", activityID)
	return nil
}
