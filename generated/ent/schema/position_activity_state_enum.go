package schema

import "entgo.io/ent/schema/field"

var PositionActivityStateEnum = field.Enum("position_activity_state").
	Values("created", "topup", "repay", "liquidation", "closed")
