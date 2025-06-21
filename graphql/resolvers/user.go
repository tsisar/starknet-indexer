package resolvers

import (
	"github.com/Tsisar/starknet-indexer/generated/ent"
	"github.com/Tsisar/starknet-indexer/generated/ent/user"
	"github.com/Tsisar/starknet-indexer/graphql/model"
)

func ApplyUserWhereInput(query *ent.UserQuery, where *model.UserWhereInput) *ent.UserQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(user.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(user.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(user.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(user.IDNotIn(where.IDNotIn...))
	}
	if where.Address != nil {
		query = query.Where(user.AddressEQ(*where.Address))
	}
	if where.AddressNot != nil {
		query = query.Where(user.AddressNEQ(*where.AddressNot))
	}
	if where.AddressIn != nil {
		query = query.Where(user.AddressIn(where.AddressIn...))
	}
	if where.AddressNotIn != nil {
		query = query.Where(user.AddressNotIn(where.AddressNotIn...))
	}
	if where.ActivePositionsCount != nil {
		query = query.Where(user.ActivePositionsCountEQ(*where.ActivePositionsCount))
	}
	if where.ActivePositionsCountNot != nil {
		query = query.Where(user.ActivePositionsCountNEQ(*where.ActivePositionsCountNot))
	}
	if where.ActivePositionsCountIn != nil {
		query = query.Where(user.ActivePositionsCountIn(where.ActivePositionsCountIn...))
	}
	if where.ActivePositionsCountNotIn != nil {
		query = query.Where(user.ActivePositionsCountNotIn(where.ActivePositionsCountNotIn...))
	}
	return query
}

func ApplyUserOrderBy(query *ent.UserQuery, orderBy *model.UserOrderBy) *ent.UserQuery {
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
	case model.UserOrderFieldID:
		query = query.Order(orderFunc(user.FieldID))
	case model.UserOrderFieldAddress:
		query = query.Order(orderFunc(user.FieldAddress))
	case model.UserOrderFieldActivePositionsCount:
		query = query.Order(orderFunc(user.FieldActivePositionsCount))
	}
	return query
}

func ApplyUserLimit(query *ent.UserQuery, first, skip *int32) *ent.UserQuery {
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
