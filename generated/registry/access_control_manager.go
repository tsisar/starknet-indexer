package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/access_control_manager"
)

var AccessControlManagerRegistry = map[string]EventMapper{
	"oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[access_control_manager.RoleGranted](nil),
	"oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[access_control_manager.RoleRevoked](nil),
	"oracle_network::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[access_control_manager.RoleAdminChanged](nil),
}
