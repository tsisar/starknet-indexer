package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/aggregator_factory"
)

var AggregatorFactoryRegistry = map[string]EventMapper{
	"oracle_network::core::aggregator_factory::AggregatorFactory::LogAggregatorClassHashUpdated": makeMapper[aggregator_factory.LogAggregatorClassHashUpdated](nil),
}
