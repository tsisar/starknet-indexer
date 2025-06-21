package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/price_feed_wbtc"
)

var PriceFeedWbtcRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":               makeMapper[price_feed_wbtc.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":             makeMapper[price_feed_wbtc.Unpaused](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetOracle":              makeMapper[price_feed_wbtc.LogSetOracle](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetTimeDelay":           makeMapper[price_feed_wbtc.LogSetTimeDelay](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogPeekPriceFailed":        makeMapper[price_feed_wbtc.LogPeekPriceFailed](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPriceLife":           makeMapper[price_feed_wbtc.LogSetPriceLife](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPoolId":              makeMapper[price_feed_wbtc.LogSetPoolId](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPrice":               makeMapper[price_feed_wbtc.LogSetPrice](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetAccessControlConfig": makeMapper[price_feed_wbtc.LogSetAccessControlConfig](nil),
}
