package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/stablecoin"
)

var StablecoinRegistry = map[string]EventMapper{
	"stablecoin::components::erc20::erc20_component::ERC20Component::Transfer":                                 makeMapper[stablecoin.Transfer](nil),
	"stablecoin::components::erc20::erc20_component::ERC20Component::Approval":                                 makeMapper[stablecoin.Approval](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[stablecoin.RoleGranted](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[stablecoin.RoleRevoked](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[stablecoin.RoleAdminChanged](nil),
}
