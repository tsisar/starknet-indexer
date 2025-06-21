package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Pool struct {
	ent.Schema
}

func (Pool) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("poolName"),
		field.String("debtCeiling"),
		field.String("liquidationRatio"),
		field.String("stabilityFeeRate"),
		field.String("tokenAdapterAddress"),
		field.String("lockedCollateral"),
		field.String("collateralPrice"),
		field.String("collateralLastPrice"),
		field.String("priceWithSafetyMargin"),
		field.String("rawPrice"),
		field.String("debtAccumulatedRate"),
		field.String("totalBorrowed"),
		field.String("totalAvailable"),
		field.String("tvl"),
	}
}

func (Pool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("positions", Position.Type),
	}
}
