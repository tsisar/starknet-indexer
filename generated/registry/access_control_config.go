package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/access_control_config"
)

var AccessControlConfigRegistry = map[string]EventMapper{
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[access_control_config.RoleGranted](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[access_control_config.RoleRevoked](nil),
	"stablecoin::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[access_control_config.RoleAdminChanged](nil),
}
