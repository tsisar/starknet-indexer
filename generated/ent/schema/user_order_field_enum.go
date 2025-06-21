package schema

import "entgo.io/ent/schema/field"

var UserOrderFieldEnum = field.Enum("user_order_field").
	Values("id", "address", "activePositionsCount")
