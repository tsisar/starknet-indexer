package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/nodes_registry"
)

var NodesRegistryRegistry = map[string]EventMapper{
	"oracle_network::core::nodes_registry::NodesRegistry::LogDeposit":                   makeMapper[nodes_registry.LogDeposit](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogNodeActivated":             makeMapper[nodes_registry.LogNodeActivated](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogNodeDeactivated":           makeMapper[nodes_registry.LogNodeDeactivated](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogWithdraw":                  makeMapper[nodes_registry.LogWithdraw](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogNodeAddedToAggregator":     makeMapper[nodes_registry.LogNodeAddedToAggregator](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogNodeRemovedFromAggregator": makeMapper[nodes_registry.LogNodeRemovedFromAggregator](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogDepositConfiscated":        makeMapper[nodes_registry.LogDepositConfiscated](nil),
	"oracle_network::core::nodes_registry::NodesRegistry::LogDepositRequiredChanged":    makeMapper[nodes_registry.LogDepositRequiredChanged](nil),
}
