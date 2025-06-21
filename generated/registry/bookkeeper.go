package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/bookkeeper"
	"github.com/tsisar/starknet-indexer/internal/stablecoin"
)

var BookkeeperRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":          makeMapper[bookkeeper.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":        makeMapper[bookkeeper.Unpaused](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogSetTotalDebtCeiling":     makeMapper[bookkeeper.LogSetTotalDebtCeiling](stablecoin.SetTotalDebtCeilingHandler),
	"stablecoin::core::bookkeeper::BookKeeper::LogSetCollateralPoolConfig": makeMapper[bookkeeper.LogSetCollateralPoolConfig](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogSetAccessControlConfig":  makeMapper[bookkeeper.LogSetAccessControlConfig](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogAdjustPosition":          makeMapper[bookkeeper.LogAdjustPosition](stablecoin.AdjustPositionHandler),
	"stablecoin::core::bookkeeper::BookKeeper::LogAddCollateral":           makeMapper[bookkeeper.LogAddCollateral](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogMoveCollateral":          makeMapper[bookkeeper.LogMoveCollateral](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogMoveStablecoin":          makeMapper[bookkeeper.LogMoveStablecoin](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogAddToWhitelist":          makeMapper[bookkeeper.LogAddToWhitelist](nil),
	"stablecoin::core::bookkeeper::BookKeeper::LogRemoveFromWhitelist":     makeMapper[bookkeeper.LogRemoveFromWhitelist](nil),
	"stablecoin::core::bookkeeper::BookKeeper::StablecoinIssuedAmount":     makeMapper[bookkeeper.StablecoinIssuedAmount](stablecoin.StablecoinIssuedAmountHandler),
	"stablecoin::core::bookkeeper::BookKeeper::LogCage":                    makeMapper[bookkeeper.LogCage](nil),
}
