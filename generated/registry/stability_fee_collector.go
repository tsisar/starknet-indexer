package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/stability_fee_collector"
)

var StabilityFeeCollectorRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                                 makeMapper[stability_fee_collector.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                               makeMapper[stability_fee_collector.Unpaused](nil),
	"stablecoin::core::stability_fee_collector::StabilityFeeCollector::LogSetSystemDebtEngine":    makeMapper[stability_fee_collector.LogSetSystemDebtEngine](nil),
	"stablecoin::core::stability_fee_collector::StabilityFeeCollector::LogNewDebtAccumulatedRate": makeMapper[stability_fee_collector.LogNewDebtAccumulatedRate](nil),
}
