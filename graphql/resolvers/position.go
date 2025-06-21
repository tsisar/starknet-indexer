package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/pool"
	"github.com/tsisar/starknet-indexer/generated/ent/position"
	"github.com/tsisar/starknet-indexer/graphql/model"
)

func ApplyPositionWhereInput(query *ent.PositionQuery, where *model.PositionWhereInput) *ent.PositionQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(position.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(position.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(position.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(position.IDNotIn(where.IDNotIn...))
	}
	if where.PositionAddress != nil {
		query = query.Where(position.PositionAddressEQ(*where.PositionAddress))
	}
	if where.PositionAddressNot != nil {
		query = query.Where(position.PositionAddressNEQ(*where.PositionAddressNot))
	}
	if where.PositionAddressIn != nil {
		query = query.Where(position.PositionAddressIn(where.PositionAddressIn...))
	}
	if where.PositionAddressNotIn != nil {
		query = query.Where(position.PositionAddressNotIn(where.PositionAddressNotIn...))
	}
	if where.UserAddress != nil {
		query = query.Where(position.UserAddressEQ(*where.UserAddress))
	}
	if where.UserAddressNot != nil {
		query = query.Where(position.UserAddressNEQ(*where.UserAddressNot))
	}
	if where.UserAddressIn != nil {
		query = query.Where(position.UserAddressIn(where.UserAddressIn...))
	}
	if where.UserAddressNotIn != nil {
		query = query.Where(position.UserAddressNotIn(where.UserAddressNotIn...))
	}
	if where.WalletAddress != nil {
		query = query.Where(position.WalletAddressEQ(*where.WalletAddress))
	}
	if where.WalletAddressNot != nil {
		query = query.Where(position.WalletAddressNEQ(*where.WalletAddressNot))
	}
	if where.WalletAddressIn != nil {
		query = query.Where(position.WalletAddressIn(where.WalletAddressIn...))
	}
	if where.WalletAddressNotIn != nil {
		query = query.Where(position.WalletAddressNotIn(where.WalletAddressNotIn...))
	}
	if where.CollateralPool != nil {
		query = query.Where(position.CollateralPoolEQ(*where.CollateralPool))
	}
	if where.CollateralPoolNot != nil {
		query = query.Where(position.CollateralPoolNEQ(*where.CollateralPoolNot))
	}
	if where.CollateralPoolIn != nil {
		query = query.Where(position.CollateralPoolIn(where.CollateralPoolIn...))
	}
	if where.CollateralPoolNotIn != nil {
		query = query.Where(position.CollateralPoolNotIn(where.CollateralPoolNotIn...))
	}
	if where.CollateralPoolName != nil {
		query = query.Where(position.CollateralPoolNameEQ(*where.CollateralPoolName))
	}
	if where.CollateralPoolNameNot != nil {
		query = query.Where(position.CollateralPoolNameNEQ(*where.CollateralPoolNameNot))
	}
	if where.CollateralPoolNameIn != nil {
		query = query.Where(position.CollateralPoolNameIn(where.CollateralPoolNameIn...))
	}
	if where.CollateralPoolNameNotIn != nil {
		query = query.Where(position.CollateralPoolNameNotIn(where.CollateralPoolNameNotIn...))
	}
	if where.PositionID != nil {
		query = query.Where(position.PositionIdEQ(*where.PositionID))
	}
	if where.PositionIDNot != nil {
		query = query.Where(position.PositionIdNEQ(*where.PositionIDNot))
	}
	if where.PositionIDIn != nil {
		query = query.Where(position.PositionIdIn(where.PositionIDIn...))
	}
	if where.PositionIDNotIn != nil {
		query = query.Where(position.PositionIdNotIn(where.PositionIDNotIn...))
	}
	if where.LockedCollateral != nil {
		query = query.Where(position.LockedCollateralEQ(*where.LockedCollateral))
	}
	if where.LockedCollateralNot != nil {
		query = query.Where(position.LockedCollateralNEQ(*where.LockedCollateralNot))
	}
	if where.LockedCollateralIn != nil {
		query = query.Where(position.LockedCollateralIn(where.LockedCollateralIn...))
	}
	if where.LockedCollateralNotIn != nil {
		query = query.Where(position.LockedCollateralNotIn(where.LockedCollateralNotIn...))
	}
	if where.DebtValue != nil {
		query = query.Where(position.DebtValueEQ(*where.DebtValue))
	}
	if where.DebtValueNot != nil {
		query = query.Where(position.DebtValueNEQ(*where.DebtValueNot))
	}
	if where.DebtValueIn != nil {
		query = query.Where(position.DebtValueIn(where.DebtValueIn...))
	}
	if where.DebtValueNotIn != nil {
		query = query.Where(position.DebtValueNotIn(where.DebtValueNotIn...))
	}
	if where.DebtShare != nil {
		query = query.Where(position.DebtShareEQ(*where.DebtShare))
	}
	if where.DebtShareNot != nil {
		query = query.Where(position.DebtShareNEQ(*where.DebtShareNot))
	}
	if where.DebtShareIn != nil {
		query = query.Where(position.DebtShareIn(where.DebtShareIn...))
	}
	if where.DebtShareNotIn != nil {
		query = query.Where(position.DebtShareNotIn(where.DebtShareNotIn...))
	}
	if where.SafetyBuffer != nil {
		query = query.Where(position.SafetyBufferEQ(*where.SafetyBuffer))
	}
	if where.SafetyBufferNot != nil {
		query = query.Where(position.SafetyBufferNEQ(*where.SafetyBufferNot))
	}
	if where.SafetyBufferIn != nil {
		query = query.Where(position.SafetyBufferIn(where.SafetyBufferIn...))
	}
	if where.SafetyBufferNotIn != nil {
		query = query.Where(position.SafetyBufferNotIn(where.SafetyBufferNotIn...))
	}
	if where.SafetyBufferInPercent != nil {
		query = query.Where(position.SafetyBufferInPercentEQ(*where.SafetyBufferInPercent))
	}
	if where.SafetyBufferInPercentNot != nil {
		query = query.Where(position.SafetyBufferInPercentNEQ(*where.SafetyBufferInPercentNot))
	}
	if where.SafetyBufferInPercentIn != nil {
		query = query.Where(position.SafetyBufferInPercentIn(where.SafetyBufferInPercentIn...))
	}
	if where.SafetyBufferInPercentNotIn != nil {
		query = query.Where(position.SafetyBufferInPercentNotIn(where.SafetyBufferInPercentNotIn...))
	}
	if where.Tvl != nil {
		query = query.Where(position.TvlEQ(*where.Tvl))
	}
	if where.TvlNot != nil {
		query = query.Where(position.TvlNEQ(*where.TvlNot))
	}
	if where.TvlIn != nil {
		query = query.Where(position.TvlIn(where.TvlIn...))
	}
	if where.TvlNotIn != nil {
		query = query.Where(position.TvlNotIn(where.TvlNotIn...))
	}
	if where.PositionStatus != nil {
		query = query.Where(position.PositionStatusEQ(string(*where.PositionStatus)))
	}
	if where.PositionStatusNot != nil {
		query = query.Where(position.PositionStatusNEQ(string(*where.PositionStatusNot)))
	}
	if where.PositionStatusIn != nil {
		statuses := make([]string, len(where.PositionStatusIn))
		for i, status := range where.PositionStatusIn {
			statuses[i] = string(status)
		}
		query = query.Where(position.PositionStatusIn(statuses...))
	}
	if where.PositionStatusNotIn != nil {
		statuses := make([]string, len(where.PositionStatusNotIn))
		for i, status := range where.PositionStatusNotIn {
			statuses[i] = string(status)
		}
		query = query.Where(position.PositionStatusNotIn(statuses...))
	}
	if where.LiquidationCount != nil {
		query = query.Where(position.LiquidationCountEQ(*where.LiquidationCount))
	}
	if where.LiquidationCountNot != nil {
		query = query.Where(position.LiquidationCountNEQ(*where.LiquidationCountNot))
	}
	if where.LiquidationCountIn != nil {
		query = query.Where(position.LiquidationCountIn(where.LiquidationCountIn...))
	}
	if where.LiquidationCountNotIn != nil {
		query = query.Where(position.LiquidationCountNotIn(where.LiquidationCountNotIn...))
	}
	if where.BlockNumber != nil {
		query = query.Where(position.BlockNumberEQ(*where.BlockNumber))
	}
	if where.BlockNumberNot != nil {
		query = query.Where(position.BlockNumberNEQ(*where.BlockNumberNot))
	}
	if where.BlockNumberIn != nil {
		query = query.Where(position.BlockNumberIn(where.BlockNumberIn...))
	}
	if where.BlockNumberNotIn != nil {
		query = query.Where(position.BlockNumberNotIn(where.BlockNumberNotIn...))
	}
	if where.BlockTimestamp != nil {
		query = query.Where(position.BlockTimestampEQ(*where.BlockTimestamp))
	}
	if where.BlockTimestampNot != nil {
		query = query.Where(position.BlockTimestampNEQ(*where.BlockTimestampNot))
	}
	if where.BlockTimestampIn != nil {
		query = query.Where(position.BlockTimestampIn(where.BlockTimestampIn...))
	}
	if where.BlockTimestampNotIn != nil {
		query = query.Where(position.BlockTimestampNotIn(where.BlockTimestampNotIn...))
	}
	if where.Transaction != nil {
		query = query.Where(position.TransactionEQ(*where.Transaction))
	}
	if where.TransactionNot != nil {
		query = query.Where(position.TransactionNEQ(*where.TransactionNot))
	}
	if where.TransactionIn != nil {
		query = query.Where(position.TransactionIn(where.TransactionIn...))
	}
	if where.TransactionNotIn != nil {
		query = query.Where(position.TransactionNotIn(where.TransactionNotIn...))
	}

	// Handle nested pool filters
	if where.Pool != nil {
		if where.Pool.PoolName != nil {
			query = query.Where(position.HasPoolWith(pool.PoolNameEQ(*where.Pool.PoolName)))
		}
		if where.Pool.PoolNameNot != nil {
			query = query.Where(position.HasPoolWith(pool.PoolNameNEQ(*where.Pool.PoolNameNot)))
		}
		if where.Pool.PoolNameIn != nil {
			query = query.Where(position.HasPoolWith(pool.PoolNameIn(where.Pool.PoolNameIn...)))
		}
		if where.Pool.PoolNameNotIn != nil {
			query = query.Where(position.HasPoolWith(pool.PoolNameNotIn(where.Pool.PoolNameNotIn...)))
		}
		if where.Pool.ID != nil {
			query = query.Where(position.HasPoolWith(pool.IDEQ(*where.Pool.ID)))
		}
		if where.Pool.IDNot != nil {
			query = query.Where(position.HasPoolWith(pool.IDNEQ(*where.Pool.IDNot)))
		}
		if where.Pool.IDIn != nil {
			query = query.Where(position.HasPoolWith(pool.IDIn(where.Pool.IDIn...)))
		}
		if where.Pool.IDNotIn != nil {
			query = query.Where(position.HasPoolWith(pool.IDNotIn(where.Pool.IDNotIn...)))
		}
	}

	return query
}

func ApplyPositionOrderBy(query *ent.PositionQuery, orderBy *model.PositionOrderBy) *ent.PositionQuery {
	if orderBy == nil {
		return query
	}

	field := orderBy.Field
	direction := orderBy.Direction

	orderFunc := ent.Asc
	if direction == model.OrderDirectionDesc {
		orderFunc = ent.Desc
	}

	switch field {
	case model.PositionOrderFieldID:
		query = query.Order(orderFunc(position.FieldID))
	case model.PositionOrderFieldPositionAddress:
		query = query.Order(orderFunc(position.FieldPositionAddress))
	case model.PositionOrderFieldUserAddress:
		query = query.Order(orderFunc(position.FieldUserAddress))
	case model.PositionOrderFieldWalletAddress:
		query = query.Order(orderFunc(position.FieldWalletAddress))
	case model.PositionOrderFieldCollateralPool:
		query = query.Order(orderFunc(position.FieldCollateralPool))
	case model.PositionOrderFieldCollateralPoolName:
		query = query.Order(orderFunc(position.FieldCollateralPoolName))
	case model.PositionOrderFieldPositionID:
		query = query.Order(orderFunc(position.FieldPositionId))
	case model.PositionOrderFieldLockedCollateral:
		query = query.Order(orderFunc(position.FieldLockedCollateral))
	case model.PositionOrderFieldDebtValue:
		query = query.Order(orderFunc(position.FieldDebtValue))
	case model.PositionOrderFieldDebtShare:
		query = query.Order(orderFunc(position.FieldDebtShare))
	case model.PositionOrderFieldSafetyBuffer:
		query = query.Order(orderFunc(position.FieldSafetyBuffer))
	case model.PositionOrderFieldSafetyBufferInPercent:
		query = query.Order(orderFunc(position.FieldSafetyBufferInPercent))
	case model.PositionOrderFieldTvl:
		query = query.Order(orderFunc(position.FieldTvl))
	case model.PositionOrderFieldPositionStatus:
		query = query.Order(orderFunc(position.FieldPositionStatus))
	case model.PositionOrderFieldLiquidationCount:
		query = query.Order(orderFunc(position.FieldLiquidationCount))
	case model.PositionOrderFieldBlockNumber:
		query = query.Order(orderFunc(position.FieldBlockNumber))
	case model.PositionOrderFieldBlockTimestamp:
		query = query.Order(orderFunc(position.FieldBlockTimestamp))
	case model.PositionOrderFieldTransaction:
		query = query.Order(orderFunc(position.FieldTransaction))
	}

	return query
}

func ApplyPositionLimit(query *ent.PositionQuery, first, skip *int32) *ent.PositionQuery {
	if skip != nil {
		query = query.Offset(int(*skip))
	}
	if first != nil {
		query = query.Limit(int(*first))
	} else {
		query = query.Limit(100)
	}
	return query
}
