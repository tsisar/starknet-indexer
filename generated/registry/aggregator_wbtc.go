package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/aggregator_wbtc"
)

var AggregatorWbtcRegistry = map[string]EventMapper{
	"oracle_network::core::aggregator::Aggregator::LogNodeAdded":       makeMapper[aggregator_wbtc.LogNodeAdded](nil),
	"oracle_network::core::aggregator::Aggregator::LogNodeRemoved":     makeMapper[aggregator_wbtc.LogNodeRemoved](nil),
	"oracle_network::core::aggregator::Aggregator::LogAnswerPublished": makeMapper[aggregator_wbtc.LogAnswerPublished](nil),
}
