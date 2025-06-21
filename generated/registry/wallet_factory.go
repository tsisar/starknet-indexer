package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/wallet_factory"
)

var WalletFactoryRegistry = map[string]EventMapper{
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[wallet_factory.RoleGranted](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[wallet_factory.RoleRevoked](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[wallet_factory.RoleAdminChanged](nil),
	"stablecoin::core::wallet::wallet_factory::WalletFactory::WalletDeployed":                                  makeMapper[wallet_factory.WalletDeployed](nil),
	"stablecoin::core::wallet::wallet_factory::WalletFactory::LogAddToWhitelist":                               makeMapper[wallet_factory.LogAddToWhitelist](nil),
	"stablecoin::core::wallet::wallet_factory::WalletFactory::LogRemoveFromWhitelist":                          makeMapper[wallet_factory.LogRemoveFromWhitelist](nil),
	"stablecoin::core::wallet::wallet_factory::WalletFactory::LogSetDecentralizedMode":                         makeMapper[wallet_factory.LogSetDecentralizedMode](nil),
}
