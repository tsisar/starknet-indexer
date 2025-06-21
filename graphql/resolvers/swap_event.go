package resolvers

import (
	"github.com/Tsisar/starknet-indexer/generated/ent"
	"github.com/Tsisar/starknet-indexer/generated/ent/swapevent"
	"github.com/Tsisar/starknet-indexer/graphql/model"
)

func ApplySwapEventWhereInput(query *ent.SwapEventQuery, where *model.SwapEventWhereInput) *ent.SwapEventQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(swapevent.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(swapevent.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(swapevent.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(swapevent.IDNotIn(where.IDNotIn...))
	}
	if where.Fee != nil {
		query = query.Where(swapevent.FeeEQ(*where.Fee))
	}
	if where.FeeNot != nil {
		query = query.Where(swapevent.FeeNEQ(*where.FeeNot))
	}
	if where.FeeIn != nil {
		query = query.Where(swapevent.FeeIn(where.FeeIn...))
	}
	if where.FeeNotIn != nil {
		query = query.Where(swapevent.FeeNotIn(where.FeeNotIn...))
	}
	if where.Owner != nil {
		query = query.Where(swapevent.OwnerEQ(*where.Owner))
	}
	if where.OwnerNot != nil {
		query = query.Where(swapevent.OwnerNEQ(*where.OwnerNot))
	}
	if where.OwnerIn != nil {
		query = query.Where(swapevent.OwnerIn(where.OwnerIn...))
	}
	if where.OwnerNotIn != nil {
		query = query.Where(swapevent.OwnerNotIn(where.OwnerNotIn...))
	}
	if where.Value != nil {
		query = query.Where(swapevent.ValueEQ(*where.Value))
	}
	if where.ValueNot != nil {
		query = query.Where(swapevent.ValueNEQ(*where.ValueNot))
	}
	if where.ValueIn != nil {
		query = query.Where(swapevent.ValueIn(where.ValueIn...))
	}
	if where.ValueNotIn != nil {
		query = query.Where(swapevent.ValueNotIn(where.ValueNotIn...))
	}
	if where.IsStablecoinToTokenSwap != nil {
		query = query.Where(swapevent.IsStablecoinToTokenSwapEQ(*where.IsStablecoinToTokenSwap))
	}
	if where.IsStablecoinToTokenSwapNot != nil {
		query = query.Where(swapevent.IsStablecoinToTokenSwapNEQ(*where.IsStablecoinToTokenSwapNot))
	}
	if where.IsTokenToStablecoinSwap != nil {
		query = query.Where(swapevent.IsTokenToStablecoinSwapEQ(*where.IsTokenToStablecoinSwap))
	}
	if where.IsTokenToStablecoinSwapNot != nil {
		query = query.Where(swapevent.IsTokenToStablecoinSwapNEQ(*where.IsTokenToStablecoinSwapNot))
	}
	if where.BlockNumber != nil {
		query = query.Where(swapevent.BlockNumberEQ(*where.BlockNumber))
	}
	if where.BlockNumberNot != nil {
		query = query.Where(swapevent.BlockNumberNEQ(*where.BlockNumberNot))
	}
	if where.BlockNumberIn != nil {
		query = query.Where(swapevent.BlockNumberIn(where.BlockNumberIn...))
	}
	if where.BlockNumberNotIn != nil {
		query = query.Where(swapevent.BlockNumberNotIn(where.BlockNumberNotIn...))
	}
	if where.BlockTimestamp != nil {
		query = query.Where(swapevent.BlockTimestampEQ(*where.BlockTimestamp))
	}
	if where.BlockTimestampNot != nil {
		query = query.Where(swapevent.BlockTimestampNEQ(*where.BlockTimestampNot))
	}
	if where.BlockTimestampIn != nil {
		query = query.Where(swapevent.BlockTimestampIn(where.BlockTimestampIn...))
	}
	if where.BlockTimestampNotIn != nil {
		query = query.Where(swapevent.BlockTimestampNotIn(where.BlockTimestampNotIn...))
	}
	if where.Transaction != nil {
		query = query.Where(swapevent.TransactionEQ(*where.Transaction))
	}
	if where.TransactionNot != nil {
		query = query.Where(swapevent.TransactionNEQ(*where.TransactionNot))
	}
	if where.TransactionIn != nil {
		query = query.Where(swapevent.TransactionIn(where.TransactionIn...))
	}
	if where.TransactionNotIn != nil {
		query = query.Where(swapevent.TransactionNotIn(where.TransactionNotIn...))
	}

	return query
}

func ApplySwapEventOrderBy(query *ent.SwapEventQuery, orderBy *model.SwapEventOrderBy) *ent.SwapEventQuery {
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
	case model.SwapEventOrderFieldID:
		query = query.Order(orderFunc(swapevent.FieldID))
	case model.SwapEventOrderFieldFee:
		query = query.Order(orderFunc(swapevent.FieldFee))
	case model.SwapEventOrderFieldOwner:
		query = query.Order(orderFunc(swapevent.FieldOwner))
	case model.SwapEventOrderFieldValue:
		query = query.Order(orderFunc(swapevent.FieldValue))
	case model.SwapEventOrderFieldIsStablecoinToTokenSwap:
		query = query.Order(orderFunc(swapevent.FieldIsStablecoinToTokenSwap))
	case model.SwapEventOrderFieldIsTokenToStablecoinSwap:
		query = query.Order(orderFunc(swapevent.FieldIsTokenToStablecoinSwap))
	case model.SwapEventOrderFieldBlockNumber:
		query = query.Order(orderFunc(swapevent.FieldBlockNumber))
	case model.SwapEventOrderFieldBlockTimestamp:
		query = query.Order(orderFunc(swapevent.FieldBlockTimestamp))
	case model.SwapEventOrderFieldTransaction:
		query = query.Order(orderFunc(swapevent.FieldTransaction))
	}
	return query
}

func ApplySwapEventLimit(query *ent.SwapEventQuery, first, skip *int32) *ent.SwapEventQuery {
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
