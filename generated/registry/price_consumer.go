package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/price_consumer"
	"github.com/Tsisar/starknet-indexer/internal/stablecoin"
)

var PriceConsumerRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                     makeMapper[price_consumer.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                   makeMapper[price_consumer.Unpaused](nil),
	"stablecoin::core::price_consumer::PriceConsumer::LogSetPrice":                    makeMapper[price_consumer.LogSetPrice](stablecoin.PriceUpdateHandler),
	"stablecoin::core::price_consumer::PriceConsumer::LogSetPriceForBatch":            makeMapper[price_consumer.LogSetPriceForBatch](nil),
	"stablecoin::core::price_consumer::PriceConsumer::LogSetStableCoinReferencePrice": makeMapper[price_consumer.LogSetStableCoinReferencePrice](nil),
	"stablecoin::core::price_consumer::PriceConsumer::LogSetBookKeeper":               makeMapper[price_consumer.LogSetBookKeeper](nil),
	"stablecoin::core::price_consumer::PriceConsumer::LogCage":                        makeMapper[price_consumer.LogCage](nil),
}
