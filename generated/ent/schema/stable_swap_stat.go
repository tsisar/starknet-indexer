package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type StableSwapStat struct {
	ent.Schema
}

func (StableSwapStat) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("totalTokenToStablecoinSwapEvents"),
		field.String("totalStablecoinToTokenSwapEvents"),
		field.String("remainingDailySwapAmount"),
		field.String("tokenToStablecoinTotalSwapValue"),
		field.String("stablecoinToTokenTotalSwapValue"),
	}
}
