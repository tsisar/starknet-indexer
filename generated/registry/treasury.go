package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/treasury"
)

var TreasuryRegistry = map[string]EventMapper{
	"oracle_network::core::treasury::Treasury::LogRewardSet":          makeMapper[treasury.LogRewardSet](nil),
	"oracle_network::core::treasury::Treasury::LogRefunded":           makeMapper[treasury.LogRefunded](nil),
	"oracle_network::core::treasury::Treasury::LogRewardsDistributed": makeMapper[treasury.LogRewardsDistributed](nil),
	"oracle_network::core::treasury::Treasury::LogFeeCollected":       makeMapper[treasury.LogFeeCollected](nil),
	"oracle_network::core::treasury::Treasury::LogFeeWithdrawn":       makeMapper[treasury.LogFeeWithdrawn](nil),
	"oracle_network::core::treasury::Treasury::LogRewardWithdrawn":    makeMapper[treasury.LogRewardWithdrawn](nil),
	"oracle_network::core::treasury::Treasury::LogSurplusWithdrawn":   makeMapper[treasury.LogSurplusWithdrawn](nil),
}
