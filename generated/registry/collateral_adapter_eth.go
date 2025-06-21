package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/collateral_adapter_eth"
)

var CollateralAdapterEthRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                                          makeMapper[collateral_adapter_eth.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                                        makeMapper[collateral_adapter_eth.Unpaused](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogDeposit":             makeMapper[collateral_adapter_eth.LogDeposit](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogWithdraw":            makeMapper[collateral_adapter_eth.LogWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogAddToWhitelist":      makeMapper[collateral_adapter_eth.LogAddToWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogRemoveFromWhitelist": makeMapper[collateral_adapter_eth.LogRemoveFromWhitelist](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogEmergencyWithdraw":   makeMapper[collateral_adapter_eth.LogEmergencyWithdraw](nil),
	"stablecoin::core::adapters::collateral_token_adapter::CollateralTokenAdapter::LogCage":                makeMapper[collateral_adapter_eth.LogCage](nil),
}
