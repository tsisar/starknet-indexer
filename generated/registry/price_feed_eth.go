package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/price_feed_eth"
)

var PriceFeedEthRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":               makeMapper[price_feed_eth.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":             makeMapper[price_feed_eth.Unpaused](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetOracle":              makeMapper[price_feed_eth.LogSetOracle](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetTimeDelay":           makeMapper[price_feed_eth.LogSetTimeDelay](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogPeekPriceFailed":        makeMapper[price_feed_eth.LogPeekPriceFailed](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPriceLife":           makeMapper[price_feed_eth.LogSetPriceLife](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPoolId":              makeMapper[price_feed_eth.LogSetPoolId](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPrice":               makeMapper[price_feed_eth.LogSetPrice](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetAccessControlConfig": makeMapper[price_feed_eth.LogSetAccessControlConfig](nil),
}
