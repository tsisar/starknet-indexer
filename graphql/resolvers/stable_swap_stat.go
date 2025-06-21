package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/stableswapstat"
	"github.com/tsisar/starknet-indexer/graphql/model"
)

func ApplyStableSwapStatWhereInput(query *ent.StableSwapStatQuery, where *model.StableSwapStatWhereInput) *ent.StableSwapStatQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(stableswapstat.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(stableswapstat.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(stableswapstat.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(stableswapstat.IDNotIn(where.IDNotIn...))
	}
	if where.TotalTokenToStablecoinSwapEvents != nil {
		query = query.Where(stableswapstat.TotalTokenToStablecoinSwapEventsEQ(*where.TotalTokenToStablecoinSwapEvents))
	}
	if where.TotalTokenToStablecoinSwapEventsNot != nil {
		query = query.Where(stableswapstat.TotalTokenToStablecoinSwapEventsNEQ(*where.TotalTokenToStablecoinSwapEventsNot))
	}
	if where.TotalTokenToStablecoinSwapEventsIn != nil {
		query = query.Where(stableswapstat.TotalTokenToStablecoinSwapEventsIn(where.TotalTokenToStablecoinSwapEventsIn...))
	}
	if where.TotalTokenToStablecoinSwapEventsNotIn != nil {
		query = query.Where(stableswapstat.TotalTokenToStablecoinSwapEventsNotIn(where.TotalTokenToStablecoinSwapEventsNotIn...))
	}
	if where.TotalStablecoinToTokenSwapEvents != nil {
		query = query.Where(stableswapstat.TotalStablecoinToTokenSwapEventsEQ(*where.TotalStablecoinToTokenSwapEvents))
	}
	if where.TotalStablecoinToTokenSwapEventsNot != nil {
		query = query.Where(stableswapstat.TotalStablecoinToTokenSwapEventsNEQ(*where.TotalStablecoinToTokenSwapEventsNot))
	}
	if where.TotalStablecoinToTokenSwapEventsIn != nil {
		query = query.Where(stableswapstat.TotalStablecoinToTokenSwapEventsIn(where.TotalStablecoinToTokenSwapEventsIn...))
	}
	if where.TotalStablecoinToTokenSwapEventsNotIn != nil {
		query = query.Where(stableswapstat.TotalStablecoinToTokenSwapEventsNotIn(where.TotalStablecoinToTokenSwapEventsNotIn...))
	}
	if where.RemainingDailySwapAmount != nil {
		query = query.Where(stableswapstat.RemainingDailySwapAmountEQ(*where.RemainingDailySwapAmount))
	}
	if where.RemainingDailySwapAmountNot != nil {
		query = query.Where(stableswapstat.RemainingDailySwapAmountNEQ(*where.RemainingDailySwapAmountNot))
	}
	if where.RemainingDailySwapAmountIn != nil {
		query = query.Where(stableswapstat.RemainingDailySwapAmountIn(where.RemainingDailySwapAmountIn...))
	}
	if where.RemainingDailySwapAmountNotIn != nil {
		query = query.Where(stableswapstat.RemainingDailySwapAmountNotIn(where.RemainingDailySwapAmountNotIn...))
	}
	if where.TokenToStablecoinTotalSwapValue != nil {
		query = query.Where(stableswapstat.TokenToStablecoinTotalSwapValueEQ(*where.TokenToStablecoinTotalSwapValue))
	}
	if where.TokenToStablecoinTotalSwapValueNot != nil {
		query = query.Where(stableswapstat.TokenToStablecoinTotalSwapValueNEQ(*where.TokenToStablecoinTotalSwapValueNot))
	}
	if where.TokenToStablecoinTotalSwapValueIn != nil {
		query = query.Where(stableswapstat.TokenToStablecoinTotalSwapValueIn(where.TokenToStablecoinTotalSwapValueIn...))
	}
	if where.TokenToStablecoinTotalSwapValueNotIn != nil {
		query = query.Where(stableswapstat.TokenToStablecoinTotalSwapValueNotIn(where.TokenToStablecoinTotalSwapValueNotIn...))
	}
	if where.StablecoinToTokenTotalSwapValue != nil {
		query = query.Where(stableswapstat.StablecoinToTokenTotalSwapValueEQ(*where.StablecoinToTokenTotalSwapValue))
	}
	if where.StablecoinToTokenTotalSwapValueNot != nil {
		query = query.Where(stableswapstat.StablecoinToTokenTotalSwapValueNEQ(*where.StablecoinToTokenTotalSwapValueNot))
	}
	if where.StablecoinToTokenTotalSwapValueIn != nil {
		query = query.Where(stableswapstat.StablecoinToTokenTotalSwapValueIn(where.StablecoinToTokenTotalSwapValueIn...))
	}
	if where.StablecoinToTokenTotalSwapValueNotIn != nil {
		query = query.Where(stableswapstat.StablecoinToTokenTotalSwapValueNotIn(where.StablecoinToTokenTotalSwapValueNotIn...))
	}
	return query
}

func ApplyStableSwapStatOrderBy(query *ent.StableSwapStatQuery, orderBy *model.StableSwapStatOrderBy) *ent.StableSwapStatQuery {
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
	case model.StableSwapStatOrderFieldID:
		query = query.Order(orderFunc(stableswapstat.FieldID))
	case model.StableSwapStatOrderFieldTotalTokenToStablecoinSwapEvents:
		query = query.Order(orderFunc(stableswapstat.FieldTotalTokenToStablecoinSwapEvents))
	case model.StableSwapStatOrderFieldTotalStablecoinToTokenSwapEvents:
		query = query.Order(orderFunc(stableswapstat.FieldTotalStablecoinToTokenSwapEvents))
	case model.StableSwapStatOrderFieldRemainingDailySwapAmount:
		query = query.Order(orderFunc(stableswapstat.FieldRemainingDailySwapAmount))
	case model.StableSwapStatOrderFieldTokenToStablecoinTotalSwapValue:
		query = query.Order(orderFunc(stableswapstat.FieldTokenToStablecoinTotalSwapValue))
	case model.StableSwapStatOrderFieldStablecoinToTokenTotalSwapValue:
		query = query.Order(orderFunc(stableswapstat.FieldStablecoinToTokenTotalSwapValue))
	}
	return query
}

func ApplyStableSwapStatLimit(query *ent.StableSwapStatQuery, first, skip *int32) *ent.StableSwapStatQuery {
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
