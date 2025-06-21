package schema

import "entgo.io/ent/schema/field"

var SwapEventOrderFieldEnum = field.Enum("swap_event_order_field").
	Values("id", "fee", "owner", "value", "isStablecoinToTokenSwap", "isTokenToStablecoinSwap", "blockNumber", "blockTimestamp", "transaction")
