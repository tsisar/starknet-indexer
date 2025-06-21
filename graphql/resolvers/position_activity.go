package resolvers

import (
	"github.com/Tsisar/starknet-indexer/generated/ent"
	"github.com/Tsisar/starknet-indexer/generated/ent/positionactivity"
	"github.com/Tsisar/starknet-indexer/graphql/model"
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
