package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/protocolstat"
	"github.com/tsisar/starknet-indexer/graphql/model"
)

func ApplyProtocolStatWhereInput(query *ent.ProtocolStatQuery, where *model.ProtocolStatWhereInput) *ent.ProtocolStatQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(protocolstat.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(protocolstat.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(protocolstat.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(protocolstat.IDNotIn(where.IDNotIn...))
	}
	if where.TotalSupply != nil {
		query = query.Where(protocolstat.TotalSupplyEQ(*where.TotalSupply))
	}
	if where.TotalSupplyNot != nil {
		query = query.Where(protocolstat.TotalSupplyNEQ(*where.TotalSupplyNot))
	}
	if where.TotalSupplyIn != nil {
		query = query.Where(protocolstat.TotalSupplyIn(where.TotalSupplyIn...))
	}
	if where.TotalSupplyNotIn != nil {
		query = query.Where(protocolstat.TotalSupplyNotIn(where.TotalSupplyNotIn...))
	}
	if where.Tvl != nil {
		query = query.Where(protocolstat.TvlEQ(*where.Tvl))
	}
	if where.TvlNot != nil {
		query = query.Where(protocolstat.TvlNEQ(*where.TvlNot))
	}
	if where.TvlIn != nil {
		query = query.Where(protocolstat.TvlIn(where.TvlIn...))
	}
	if where.TvlNotIn != nil {
		query = query.Where(protocolstat.TvlNotIn(where.TvlNotIn...))
	}
	return query
}

func ApplyProtocolStatOrderBy(query *ent.ProtocolStatQuery, orderBy *model.ProtocolStatOrderBy) *ent.ProtocolStatQuery {
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
	case model.ProtocolStatOrderFieldID:
		query = query.Order(orderFunc(protocolstat.FieldID))
	case model.ProtocolStatOrderFieldTotalSupply:
		query = query.Order(orderFunc(protocolstat.FieldTotalSupply))
	case model.ProtocolStatOrderFieldTvl:
		query = query.Order(orderFunc(protocolstat.FieldTvl))
	}

	return query
}

func ApplyProtocolStatLimit(query *ent.ProtocolStatQuery, first, skip *int32) *ent.ProtocolStatQuery {
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
