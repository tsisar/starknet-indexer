package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Position struct {
	ent.Schema
}

func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("positionAddress"),
		field.String("userAddress"),
		field.String("walletAddress"),
		field.String("collateralPool"),
		field.String("collateralPoolName"),
		field.String("positionId"),
		field.String("lockedCollateral"),
		field.String("debtValue"),
		field.String("debtShare"),
		field.String("safetyBuffer"),
		field.String("safetyBufferInPercent"),
		field.String("tvl"),
		field.String("positionStatus"),
		field.String("liquidationCount"),
		field.String("blockNumber"),
		field.String("blockTimestamp"),
		field.String("transaction"),
	}
}

func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pool", Pool.Type).Unique(),
		edge.To("activity", PositionActivity.Type),
	}
}
