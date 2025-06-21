package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PositionActivity struct {
	ent.Schema
}

func (PositionActivity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("activityState"),
		field.String("collateralAmount"),
		field.String("debtAmount"),
		field.String("blockNumber"),
		field.String("blockTimestamp"),
		field.String("transaction"),
	}
}

func (PositionActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("position", Position.Type).Unique(),
	}
}
