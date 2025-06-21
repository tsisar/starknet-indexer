package registry

import (
	"github.com/Tsisar/starknet-indexer/generated/subgraph/position_manager"
	"github.com/Tsisar/starknet-indexer/internal/stablecoin"
)

var PositionManagerRegistry = map[string]EventMapper{
	"stablecoin::components::pausable::PausableComponent::Paused":                            makeMapper[position_manager.Paused](nil),
	"stablecoin::components::pausable::PausableComponent::Unpaused":                          makeMapper[position_manager.Unpaused](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogNewPosition":          makeMapper[position_manager.LogNewPosition](stablecoin.NewPositionHandler),
	"stablecoin::core::position::position_manager::PositionManager::LogAllowManagePosition":  makeMapper[position_manager.LogAllowManagePosition](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogAllowMigratePosition": makeMapper[position_manager.LogAllowMigratePosition](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogExportPosition":       makeMapper[position_manager.LogExportPosition](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogImportPosition":       makeMapper[position_manager.LogImportPosition](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogMovePosition":         makeMapper[position_manager.LogMovePosition](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogPriceConsumerUpdated": makeMapper[position_manager.LogPriceConsumerUpdated](nil),
	"stablecoin::core::position::position_manager::PositionManager::LogBookKeeperUpdated":    makeMapper[position_manager.LogBookKeeperUpdated](nil),
}
