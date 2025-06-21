package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProtocolStat struct {
	ent.Schema
}

func (ProtocolStat) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("totalSupply"),
		field.String("tvl"),
		field.String("pools"),
	}
}
