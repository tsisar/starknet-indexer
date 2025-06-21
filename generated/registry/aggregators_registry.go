package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/aggregators_registry"
)

var AggregatorsRegistryRegistry = map[string]EventMapper{
	"oracle_network::core::aggregators_registry::AggregatorsRegistry::LogAggregatorCreated":        makeMapper[aggregators_registry.LogAggregatorCreated](nil),
	"oracle_network::core::aggregators_registry::AggregatorsRegistry::LogSubscriptionPriceChanged": makeMapper[aggregators_registry.LogSubscriptionPriceChanged](nil),
}
