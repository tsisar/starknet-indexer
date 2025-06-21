package schema

import "entgo.io/ent/schema/field"

var OrderDirectionEnum = field.Enum("order_direction").
	Values("asc", "desc")
