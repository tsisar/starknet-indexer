package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/splyce_price_oracle_wbtc"
)

var SplycePriceOracleWbtcRegistry = map[string]EventMapper{
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[splyce_price_oracle_wbtc.RoleGranted](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[splyce_price_oracle_wbtc.RoleRevoked](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[splyce_price_oracle_wbtc.RoleAdminChanged](nil),
	"stablecoin::components::pausable::PausableComponent::Paused":                                              makeMapper[splyce_price_oracle_wbtc.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                                            makeMapper[splyce_price_oracle_wbtc.Unpaused](nil),
	"stablecoin::core::price::splyce_price_oracle::SplycePriceOracle::LogSetOracleAggregator":                  makeMapper[splyce_price_oracle_wbtc.LogSetOracleAggregator](nil),
}
