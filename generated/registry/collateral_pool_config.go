package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/collateral_pool_config"
	"github.com/Tsisar/starknet-indexer/internal/stablecoin"
)

var CollateralPoolConfigRegistry = map[string]EventMapper{
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetPriceWithSafetyMargin":  makeMapper[collateral_pool_config.LogSetPriceWithSafetyMargin](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetDebtCeiling":            makeMapper[collateral_pool_config.LogSetDebtCeiling](stablecoin.HandleLogSetDebtCeiling),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetDebtFloor":              makeMapper[collateral_pool_config.LogSetDebtFloor](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetPriceFeed":              makeMapper[collateral_pool_config.LogSetPriceFeed](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetLiquidationRatio":       makeMapper[collateral_pool_config.LogSetLiquidationRatio](stablecoin.HandleSetLiquidationRatio),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetStabilityFeeRate":       makeMapper[collateral_pool_config.LogSetStabilityFeeRate](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetAdapter":                makeMapper[collateral_pool_config.LogSetAdapter](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetCloseFactorBps":         makeMapper[collateral_pool_config.LogSetCloseFactorBps](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetLiquidatorIncentiveBps": makeMapper[collateral_pool_config.LogSetLiquidatorIncentiveBps](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetTreasuryFeesBps":        makeMapper[collateral_pool_config.LogSetTreasuryFeesBps](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetStrategy":               makeMapper[collateral_pool_config.LogSetStrategy](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetTotalDebtShare":         makeMapper[collateral_pool_config.LogSetTotalDebtShare](nil),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogSetDebtAccumulatedRate":    makeMapper[collateral_pool_config.LogSetDebtAccumulatedRate](stablecoin.HandleSetDebtAccumulatedRate),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogInitCollateralPoolId":      makeMapper[collateral_pool_config.LogInitCollateralPoolId](stablecoin.HandleLogInitCollateralPoolId),
	"stablecoin::core::config::collateral_pool_config::CollateralPoolConfig::LogPositionDebtCeiling":       makeMapper[collateral_pool_config.LogPositionDebtCeiling](nil),
}
