package contract

import (
	"context"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/internal/config"
	helper "github.com/tsisar/starknet-indexer/internal/stablecoin/utils"
	"math/big"
)

func FetchDebtAccumulatedRate(ctx context.Context, poolId string) *big.Float {
	log.Infof("Fetching debt accumulated rate for pool ID: %s", poolId)

	contractAddress := config.App.Contracts["COLLATERAL_POOL_CONFIG"]
	log.Debugf("Using COLLATERAL_POOL_CONFIG address: %s", contractAddress)

	feltPoolId, err := utils.HexToFelt(poolId)
	if err != nil {
		log.Errorf("Failed to convert pool ID to Felt: %v", err)
		return nil
	}

	result, err := callFunction(ctx, contractAddress, "get_debt_accumulated_rate", feltPoolId)
	if err != nil {
		log.Errorf("Failed to call get_debt_accumulated_rate function: %v", err)
		return nil
	}

	log.Debugf("Raw result from get_debt_accumulated_rate: %v", result)

	rate := utils.FeltToBigInt(result[0])
	finalRate := helper.DivByRAYToDecimal(rate)
	log.Infof("Calculated debt accumulated rate: %v", finalRate)

	return finalRate
}
