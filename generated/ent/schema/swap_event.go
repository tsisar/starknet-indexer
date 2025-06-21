package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SwapEvent struct {
	ent.Schema
}

func (SwapEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("fee"),
		field.String("owner"),
		field.String("value"),
		field.Bool("isStablecoinToTokenSwap"),
		field.Bool("isTokenToStablecoinSwap"),
		field.String("blockNumber"),
		field.String("blockTimestamp"),
		field.String("transaction"),
	}
}
