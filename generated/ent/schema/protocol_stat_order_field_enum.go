package schema

import "entgo.io/ent/schema/field"

var ProtocolStatOrderFieldEnum = field.Enum("protocol_stat_order_field").
	Values("id", "totalSupply", "tvl")
