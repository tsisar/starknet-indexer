package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/factory"
)

var FactoryRegistry = map[string]EventMapper{
	"vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleGranted":      makeMapper[factory.RoleGranted](nil),
	"vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleRevoked":      makeMapper[factory.RoleRevoked](nil),
	"vaults::external::components::accesscontrol::accesscontrol_component::AccessControlComponent::RoleAdminChanged": makeMapper[factory.RoleAdminChanged](nil),
	"vaults::factory::Factory::FeeConfigUpdated":                                                                     makeMapper[factory.FeeConfigUpdated](nil),
	"vaults::factory::Factory::VaultDeployed":                                                                        makeMapper[factory.VaultDeployed](nil),
	"vaults::factory::Factory::VaultClassHashUpdated":                                                                makeMapper[factory.VaultClassHashUpdated](nil),
}
