package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/liquidation_strategy"
	"github.com/tsisar/starknet-indexer/internal/stablecoin"
)

var LiquidationStrategyRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                                                                makeMapper[liquidation_strategy.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                                                              makeMapper[liquidation_strategy.Unpaused](nil),
	"stablecoin::core::strategies::fixed_spread_liquidation_strategy::FixedSpreadLiquidationStrategy::LogFixedSpreadLiquidate":   makeMapper[liquidation_strategy.LogFixedSpreadLiquidate](stablecoin.PositionLiquidationHandler),
	"stablecoin::core::strategies::fixed_spread_liquidation_strategy::FixedSpreadLiquidationStrategy::LogSetFlashLendingEnabled": makeMapper[liquidation_strategy.LogSetFlashLendingEnabled](nil),
	"stablecoin::core::strategies::fixed_spread_liquidation_strategy::FixedSpreadLiquidationStrategy::LogSetBookKeeper":          makeMapper[liquidation_strategy.LogSetBookKeeper](nil),
}
