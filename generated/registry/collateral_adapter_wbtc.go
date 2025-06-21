package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/collateral_adapter_wbtc"
)

var CollateralAdapterWbtcRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                                          makeMapper[collateral_adapter_wbtc.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                                        makeMapper[collateral_adapter_wbtc.Unpaused](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogDeposit":             makeMapper[collateral_adapter_wbtc.LogDeposit](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogWithdraw":            makeMapper[collateral_adapter_wbtc.LogWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogAddToWhitelist":      makeMapper[collateral_adapter_wbtc.LogAddToWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogRemoveFromWhitelist": makeMapper[collateral_adapter_wbtc.LogRemoveFromWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogEmergencyWithdraw":   makeMapper[collateral_adapter_wbtc.LogEmergencyWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogCage":                makeMapper[collateral_adapter_wbtc.LogCage](nil),
}
