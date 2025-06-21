package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/liquidation_engine"
)

var LiquidationEngineRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                     makeMapper[liquidation_engine.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                   makeMapper[liquidation_engine.Unpaused](nil),
	"stablecoin::core::liquidation_engine::LiquidationEngine::LiquidationFail":        makeMapper[liquidation_engine.LiquidationFail](nil),
	"stablecoin::core::liquidation_engine::LiquidationEngine::LogSetBookKeeper":       makeMapper[liquidation_engine.LogSetBookKeeper](nil),
	"stablecoin::core::liquidation_engine::LiquidationEngine::LogAddToWhitelist":      makeMapper[liquidation_engine.LogAddToWhitelist](nil),
	"stablecoin::core::liquidation_engine::LiquidationEngine::LogRemoveFromWhitelist": makeMapper[liquidation_engine.LogRemoveFromWhitelist](nil),
	"stablecoin::core::liquidation_engine::LiquidationEngine::LogCage":                makeMapper[liquidation_engine.LogCage](nil),
}
