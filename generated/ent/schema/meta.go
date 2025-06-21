package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Meta struct {
	ent.Schema
}

func (Meta) Fields() []ent.Field {
	return []ent.Field{
		field.String("deployment"),
		field.String("hasIndexingErrors"),
	}
}

func (Meta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("block", Block.Type).Unique(),
	}
}
