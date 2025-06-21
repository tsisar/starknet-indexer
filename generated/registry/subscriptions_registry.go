package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/subscriptions_registry"
)

var SubscriptionsRegistryRegistry = map[string]EventMapper{
	"oracle_network::core::subscriptions_registry::SubscriptionsRegistry::LogSubscribed":           makeMapper[subscriptions_registry.LogSubscribed](nil),
	"oracle_network::core::subscriptions_registry::SubscriptionsRegistry::LogUnsubscribed":         makeMapper[subscriptions_registry.LogUnsubscribed](nil),
	"oracle_network::core::subscriptions_registry::SubscriptionsRegistry::LogSubscriptionPriceSet": makeMapper[subscriptions_registry.LogSubscriptionPriceSet](nil),
	"oracle_network::core::subscriptions_registry::SubscriptionsRegistry::LogSubscriptionFeeSet":   makeMapper[subscriptions_registry.LogSubscriptionFeeSet](nil),
}
