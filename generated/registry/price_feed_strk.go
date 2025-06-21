package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/price_feed_strk"
)

var PriceFeedStrkRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":               makeMapper[price_feed_strk.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":             makeMapper[price_feed_strk.Unpaused](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetOracle":              makeMapper[price_feed_strk.LogSetOracle](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetTimeDelay":           makeMapper[price_feed_strk.LogSetTimeDelay](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogPeekPriceFailed":        makeMapper[price_feed_strk.LogPeekPriceFailed](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPriceLife":           makeMapper[price_feed_strk.LogSetPriceLife](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPoolId":              makeMapper[price_feed_strk.LogSetPoolId](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetPrice":               makeMapper[price_feed_strk.LogSetPrice](nil),
	"stablecoin::core::price::price_feed::PriceFeed::LogSetAccessControlConfig": makeMapper[price_feed_strk.LogSetAccessControlConfig](nil),
}
