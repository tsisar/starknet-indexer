package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/show_stopper"
)

var ShowStopperRegistry = map[string]EventMapper{
	"stablecoin::core::show_stopper::ShowStopper::LogCage":                   makeMapper[show_stopper.LogCage](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogCageCollateralPool":     makeMapper[show_stopper.LogCageCollateralPool](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogAccumulateBadDebt":      makeMapper[show_stopper.LogAccumulateBadDebt](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogRedeemLockedCollateral": makeMapper[show_stopper.LogRedeemLockedCollateral](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogFinalizeDebt":           makeMapper[show_stopper.LogFinalizeDebt](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogFinalizeCashPrice":      makeMapper[show_stopper.LogFinalizeCashPrice](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogAccumulateStablecoin":   makeMapper[show_stopper.LogAccumulateStablecoin](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogRedeemStablecoin":       makeMapper[show_stopper.LogRedeemStablecoin](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogSetBookKeeper":          makeMapper[show_stopper.LogSetBookKeeper](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogSetLiquidationEngine":   makeMapper[show_stopper.LogSetLiquidationEngine](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogSetSystemDebtEngine":    makeMapper[show_stopper.LogSetSystemDebtEngine](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogSetPriceConsumer":       makeMapper[show_stopper.LogSetPriceConsumer](nil),
	"stablecoin::core::show_stopper::ShowStopper::LogSetCageCoolDown":        makeMapper[show_stopper.LogSetCageCoolDown](nil),
}
