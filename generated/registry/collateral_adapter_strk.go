package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/collateral_adapter_strk"
)

var CollateralAdapterStrkRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                                          makeMapper[collateral_adapter_strk.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                                        makeMapper[collateral_adapter_strk.Unpaused](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogDeposit":             makeMapper[collateral_adapter_strk.LogDeposit](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogWithdraw":            makeMapper[collateral_adapter_strk.LogWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogAddToWhitelist":      makeMapper[collateral_adapter_strk.LogAddToWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogRemoveFromWhitelist": makeMapper[collateral_adapter_strk.LogRemoveFromWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogEmergencyWithdraw":   makeMapper[collateral_adapter_strk.LogEmergencyWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogCage":                makeMapper[collateral_adapter_strk.LogCage](nil),
}
