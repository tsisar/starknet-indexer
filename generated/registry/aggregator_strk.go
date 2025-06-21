package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/aggregator_strk"
)

var AggregatorStrkRegistry = map[string]EventMapper{
	"oracle_network::core::aggregator::Aggregator::LogNodeAdded":       makeMapper[aggregator_strk.LogNodeAdded](nil),
	"oracle_network::core::aggregator::Aggregator::LogNodeRemoved":     makeMapper[aggregator_strk.LogNodeRemoved](nil),
	"oracle_network::core::aggregator::Aggregator::LogAnswerPublished": makeMapper[aggregator_strk.LogAnswerPublished](nil),
}
