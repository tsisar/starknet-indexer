package registry

import (
	"github.com/tsisar/starknet-indexer/generated/subgraph/admin_controls"
)

var AdminControlsRegistry = map[string]EventMapper{
	"stablecoin::core::admin_controls::AdminControls::LogPauseProtocol":        makeMapper[admin_controls.LogPauseProtocol](nil),
	"stablecoin::core::admin_controls::AdminControls::LogUnpauseProtocol":      makeMapper[admin_controls.LogUnpauseProtocol](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetBookKeeper":        makeMapper[admin_controls.LogSetBookKeeper](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetLiquidationEngine": makeMapper[admin_controls.LogSetLiquidationEngine](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetPriceConsumer":     makeMapper[admin_controls.LogSetPriceConsumer](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetPositionManager":   makeMapper[admin_controls.LogSetPositionManager](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetSystemDebtEngine":  makeMapper[admin_controls.LogSetSystemDebtEngine](nil),
	"stablecoin::core::admin_controls::AdminControls::LogSetStablecoinAdapter": makeMapper[admin_controls.LogSetStablecoinAdapter](nil),
}
