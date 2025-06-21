package schema

import "entgo.io/ent/schema/field"

var PositionStatusEnum = field.Enum("position_status").
	Values("safe", "unsafe", "closed")
