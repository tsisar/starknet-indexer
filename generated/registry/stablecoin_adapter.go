package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/stablecoin_adapter"
)

var StablecoinAdapterRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                makeMapper[stablecoin_adapter.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":              makeMapper[stablecoin_adapter.Unpaused](nil),
	"stablecoin::core::adapters::stablecoin_adapter::StablecoinAdapter::LogCage": makeMapper[stablecoin_adapter.LogCage](nil),
}
