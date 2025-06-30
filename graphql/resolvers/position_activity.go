package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/position"
	"github.com/tsisar/starknet-indexer/generated/ent/positionactivity"
	"github.com/tsisar/starknet-indexer/graphql/model"
)

func ApplyPositionActivityWhereInput(query *ent.PositionActivityQuery, where *model.PositionActivityWhereInput) *ent.PositionActivityQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(positionactivity.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(positionactivity.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(positionactivity.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(positionactivity.IDNotIn(where.IDNotIn...))
	}
	if where.ActivityState != nil {
		query = query.Where(positionactivity.ActivityStateEQ(string(*where.ActivityState)))
	}
	if where.ActivityStateNot != nil {
		query = query.Where(positionactivity.ActivityStateNEQ(string(*where.ActivityStateNot)))
	}
	if where.ActivityStateIn != nil {
		states := make([]string, len(where.ActivityStateIn))
		for i, state := range where.ActivityStateIn {
			states[i] = string(state)
		}
		query = query.Where(positionactivity.ActivityStateIn(states...))
	}
	if where.ActivityStateNotIn != nil {
		states := make([]string, len(where.ActivityStateNotIn))
		for i, state := range where.ActivityStateNotIn {
			states[i] = string(state)
		}
		query = query.Where(positionactivity.ActivityStateNotIn(states...))
	}
	if where.CollateralAmount != nil {
		query = query.Where(positionactivity.CollateralAmountEQ(*where.CollateralAmount))
	}
	if where.CollateralAmountNot != nil {
		query = query.Where(positionactivity.CollateralAmountNEQ(*where.CollateralAmountNot))
	}
	if where.CollateralAmountIn != nil {
		query = query.Where(positionactivity.CollateralAmountIn(where.CollateralAmountIn...))
	}
	if where.CollateralAmountNotIn != nil {
		query = query.Where(positionactivity.CollateralAmountNotIn(where.CollateralAmountNotIn...))
	}
	if where.DebtAmount != nil {
		query = query.Where(positionactivity.DebtAmountEQ(*where.DebtAmount))
	}
	if where.DebtAmountNot != nil {
		query = query.Where(positionactivity.DebtAmountNEQ(*where.DebtAmountNot))
	}
	if where.DebtAmountIn != nil {
		query = query.Where(positionactivity.DebtAmountIn(where.DebtAmountIn...))
	}
	if where.DebtAmountNotIn != nil {
		query = query.Where(positionactivity.DebtAmountNotIn(where.DebtAmountNotIn...))
	}
	if where.BlockNumber != nil {
		query = query.Where(positionactivity.BlockNumberEQ(*where.BlockNumber))
	}
	if where.BlockNumberNot != nil {
		query = query.Where(positionactivity.BlockNumberNEQ(*where.BlockNumberNot))
	}
	if where.BlockNumberIn != nil {
		query = query.Where(positionactivity.BlockNumberIn(where.BlockNumberIn...))
	}
	if where.BlockNumberNotIn != nil {
		query = query.Where(positionactivity.BlockNumberNotIn(where.BlockNumberNotIn...))
	}
	if where.BlockTimestamp != nil {
		query = query.Where(positionactivity.BlockTimestampEQ(*where.BlockTimestamp))
	}
	if where.BlockTimestampNot != nil {
		query = query.Where(positionactivity.BlockTimestampNEQ(*where.BlockTimestampNot))
	}
	if where.BlockTimestampIn != nil {
		query = query.Where(positionactivity.BlockTimestampIn(where.BlockTimestampIn...))
	}
	if where.BlockTimestampNotIn != nil {
		query = query.Where(positionactivity.BlockTimestampNotIn(where.BlockTimestampNotIn...))
	}
	if where.Transaction != nil {
		query = query.Where(positionactivity.TransactionEQ(*where.Transaction))
	}
	if where.TransactionNot != nil {
		query = query.Where(positionactivity.TransactionNEQ(*where.TransactionNot))
	}
	if where.TransactionIn != nil {
		query = query.Where(positionactivity.TransactionIn(where.TransactionIn...))
	}
	if where.TransactionNotIn != nil {
		query = query.Where(positionactivity.TransactionNotIn(where.TransactionNotIn...))
	}

	// Handle nested position filters
	if where.Position != nil {
		if where.Position.ID != nil {
			query = query.Where(positionactivity.HasPositionWith(position.IDEQ(*where.Position.ID)))
		}
		if where.Position.IDNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.IDNEQ(*where.Position.IDNot)))
		}
		if where.Position.IDIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.IDIn(where.Position.IDIn...)))
		}
		if where.Position.IDNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.IDNotIn(where.Position.IDNotIn...)))
		}
		if where.Position.PositionAddress != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionAddressEQ(*where.Position.PositionAddress)))
		}
		if where.Position.PositionAddressNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionAddressNEQ(*where.Position.PositionAddressNot)))
		}
		if where.Position.PositionAddressIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionAddressIn(where.Position.PositionAddressIn...)))
		}
		if where.Position.PositionAddressNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionAddressNotIn(where.Position.PositionAddressNotIn...)))
		}
		if where.Position.UserAddress != nil {
			query = query.Where(positionactivity.HasPositionWith(position.UserAddressEQ(*where.Position.UserAddress)))
		}
		if where.Position.UserAddressNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.UserAddressNEQ(*where.Position.UserAddressNot)))
		}
		if where.Position.UserAddressIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.UserAddressIn(where.Position.UserAddressIn...)))
		}
		if where.Position.UserAddressNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.UserAddressNotIn(where.Position.UserAddressNotIn...)))
		}
		if where.Position.WalletAddress != nil {
			query = query.Where(positionactivity.HasPositionWith(position.WalletAddressEQ(*where.Position.WalletAddress)))
		}
		if where.Position.WalletAddressNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.WalletAddressNEQ(*where.Position.WalletAddressNot)))
		}
		if where.Position.WalletAddressIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.WalletAddressIn(where.Position.WalletAddressIn...)))
		}
		if where.Position.WalletAddressNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.WalletAddressNotIn(where.Position.WalletAddressNotIn...)))
		}
		if where.Position.CollateralPool != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolEQ(*where.Position.CollateralPool)))
		}
		if where.Position.CollateralPoolNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNEQ(*where.Position.CollateralPoolNot)))
		}
		if where.Position.CollateralPoolIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolIn(where.Position.CollateralPoolIn...)))
		}
		if where.Position.CollateralPoolNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNotIn(where.Position.CollateralPoolNotIn...)))
		}
		if where.Position.CollateralPoolName != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNameEQ(*where.Position.CollateralPoolName)))
		}
		if where.Position.CollateralPoolNameNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNameNEQ(*where.Position.CollateralPoolNameNot)))
		}
		if where.Position.CollateralPoolNameIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNameIn(where.Position.CollateralPoolNameIn...)))
		}
		if where.Position.CollateralPoolNameNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.CollateralPoolNameNotIn(where.Position.CollateralPoolNameNotIn...)))
		}
		if where.Position.PositionID != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionIdEQ(*where.Position.PositionID)))
		}
		if where.Position.PositionIDNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionIdNEQ(*where.Position.PositionIDNot)))
		}
		if where.Position.PositionIDIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionIdIn(where.Position.PositionIDIn...)))
		}
		if where.Position.PositionIDNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionIdNotIn(where.Position.PositionIDNotIn...)))
		}
		if where.Position.LockedCollateral != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LockedCollateralEQ(*where.Position.LockedCollateral)))
		}
		if where.Position.LockedCollateralNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LockedCollateralNEQ(*where.Position.LockedCollateralNot)))
		}
		if where.Position.LockedCollateralIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LockedCollateralIn(where.Position.LockedCollateralIn...)))
		}
		if where.Position.LockedCollateralNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LockedCollateralNotIn(where.Position.LockedCollateralNotIn...)))
		}
		if where.Position.DebtValue != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtValueEQ(*where.Position.DebtValue)))
		}
		if where.Position.DebtValueNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtValueNEQ(*where.Position.DebtValueNot)))
		}
		if where.Position.DebtValueIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtValueIn(where.Position.DebtValueIn...)))
		}
		if where.Position.DebtValueNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtValueNotIn(where.Position.DebtValueNotIn...)))
		}
		if where.Position.DebtShare != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtShareEQ(*where.Position.DebtShare)))
		}
		if where.Position.DebtShareNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtShareNEQ(*where.Position.DebtShareNot)))
		}
		if where.Position.DebtShareIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtShareIn(where.Position.DebtShareIn...)))
		}
		if where.Position.DebtShareNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.DebtShareNotIn(where.Position.DebtShareNotIn...)))
		}
		if where.Position.SafetyBuffer != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferEQ(*where.Position.SafetyBuffer)))
		}
		if where.Position.SafetyBufferNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferNEQ(*where.Position.SafetyBufferNot)))
		}
		if where.Position.SafetyBufferIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferIn(where.Position.SafetyBufferIn...)))
		}
		if where.Position.SafetyBufferNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferNotIn(where.Position.SafetyBufferNotIn...)))
		}
		if where.Position.SafetyBufferInPercent != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferInPercentEQ(*where.Position.SafetyBufferInPercent)))
		}
		if where.Position.SafetyBufferInPercentNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferInPercentNEQ(*where.Position.SafetyBufferInPercentNot)))
		}
		if where.Position.SafetyBufferInPercentIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferInPercentIn(where.Position.SafetyBufferInPercentIn...)))
		}
		if where.Position.SafetyBufferInPercentNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.SafetyBufferInPercentNotIn(where.Position.SafetyBufferInPercentNotIn...)))
		}
		if where.Position.Tvl != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TvlEQ(*where.Position.Tvl)))
		}
		if where.Position.TvlNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TvlNEQ(*where.Position.TvlNot)))
		}
		if where.Position.TvlIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TvlIn(where.Position.TvlIn...)))
		}
		if where.Position.TvlNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TvlNotIn(where.Position.TvlNotIn...)))
		}
		if where.Position.PositionStatus != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionStatusEQ(string(*where.Position.PositionStatus))))
		}
		if where.Position.PositionStatusNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.PositionStatusNEQ(string(*where.Position.PositionStatusNot))))
		}
		if where.Position.PositionStatusIn != nil {
			statuses := make([]string, len(where.Position.PositionStatusIn))
			for i, status := range where.Position.PositionStatusIn {
				statuses[i] = string(status)
			}
			query = query.Where(positionactivity.HasPositionWith(position.PositionStatusIn(statuses...)))
		}
		if where.Position.PositionStatusNotIn != nil {
			statuses := make([]string, len(where.Position.PositionStatusNotIn))
			for i, status := range where.Position.PositionStatusNotIn {
				statuses[i] = string(status)
			}
			query = query.Where(positionactivity.HasPositionWith(position.PositionStatusNotIn(statuses...)))
		}
		if where.Position.LiquidationCount != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LiquidationCountEQ(*where.Position.LiquidationCount)))
		}
		if where.Position.LiquidationCountNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LiquidationCountNEQ(*where.Position.LiquidationCountNot)))
		}
		if where.Position.LiquidationCountIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LiquidationCountIn(where.Position.LiquidationCountIn...)))
		}
		if where.Position.LiquidationCountNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.LiquidationCountNotIn(where.Position.LiquidationCountNotIn...)))
		}
		if where.Position.BlockNumber != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockNumberEQ(*where.Position.BlockNumber)))
		}
		if where.Position.BlockNumberNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockNumberNEQ(*where.Position.BlockNumberNot)))
		}
		if where.Position.BlockNumberIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockNumberIn(where.Position.BlockNumberIn...)))
		}
		if where.Position.BlockNumberNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockNumberNotIn(where.Position.BlockNumberNotIn...)))
		}
		if where.Position.BlockTimestamp != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockTimestampEQ(*where.Position.BlockTimestamp)))
		}
		if where.Position.BlockTimestampNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockTimestampNEQ(*where.Position.BlockTimestampNot)))
		}
		if where.Position.BlockTimestampIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockTimestampIn(where.Position.BlockTimestampIn...)))
		}
		if where.Position.BlockTimestampNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.BlockTimestampNotIn(where.Position.BlockTimestampNotIn...)))
		}
		if where.Position.Transaction != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TransactionEQ(*where.Position.Transaction)))
		}
		if where.Position.TransactionNot != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TransactionNEQ(*where.Position.TransactionNot)))
		}
		if where.Position.TransactionIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TransactionIn(where.Position.TransactionIn...)))
		}
		if where.Position.TransactionNotIn != nil {
			query = query.Where(positionactivity.HasPositionWith(position.TransactionNotIn(where.Position.TransactionNotIn...)))
		}
	}

	return query
}

func ApplyPositionActivityOrderBy(query *ent.PositionActivityQuery, orderBy *model.PositionActivityOrderBy) *ent.PositionActivityQuery {
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
	case model.PositionActivityOrderFieldID:
		query = query.Order(orderFunc(positionactivity.FieldID))
	case model.PositionActivityOrderFieldActivityState:
		query = query.Order(orderFunc(positionactivity.FieldActivityState))
	case model.PositionActivityOrderFieldCollateralAmount:
		query = query.Order(orderFunc(positionactivity.FieldCollateralAmount))
	case model.PositionActivityOrderFieldDebtAmount:
		query = query.Order(orderFunc(positionactivity.FieldDebtAmount))
	case model.PositionActivityOrderFieldBlockNumber:
		query = query.Order(orderFunc(positionactivity.FieldBlockNumber))
	case model.PositionActivityOrderFieldBlockTimestamp:
		query = query.Order(orderFunc(positionactivity.FieldBlockTimestamp))
	case model.PositionActivityOrderFieldTransaction:
		query = query.Order(orderFunc(positionactivity.FieldTransaction))
	}
	return query
}

func ApplyPositionActivityLimit(query *ent.PositionActivityQuery, first, skip *int32) *ent.PositionActivityQuery {
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
