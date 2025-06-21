package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/system_debt_engine"
)

var SystemDebtEngineRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                 makeMapper[system_debt_engine.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":               makeMapper[system_debt_engine.Unpaused](nil),
	"stablecoin::core::system_debt_engine::SystemDebtEngine::LogSetSurplusBuffer": makeMapper[system_debt_engine.LogSetSurplusBuffer](nil),
	"stablecoin::core::system_debt_engine::SystemDebtEngine::LogCage":             makeMapper[system_debt_engine.LogCage](nil),
}
