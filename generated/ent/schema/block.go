package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Block struct {
	ent.Schema
}

func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash"),
		field.String("number"),
		field.String("parentHash"),
		field.String("timestamp"),
	}
}
