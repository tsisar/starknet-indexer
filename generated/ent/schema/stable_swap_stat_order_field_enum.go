package schema

import "entgo.io/ent/schema/field"

var StableSwapStatOrderFieldEnum = field.Enum("stable_swap_stat_order_field").
	Values("id", "totalTokenToStablecoinSwapEvents", "totalStablecoinToTokenSwapEvents", "remainingDailySwapAmount", "tokenToStablecoinTotalSwapValue", "stablecoinToTokenTotalSwapValue")
