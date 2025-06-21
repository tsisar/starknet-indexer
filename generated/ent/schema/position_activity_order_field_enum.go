package schema

import "entgo.io/ent/schema/field"

var PositionActivityOrderFieldEnum = field.Enum("position_activity_order_field").
	Values("id", "position", "activityState", "collateralAmount", "debtAmount", "blockNumber", "blockTimestamp", "transaction")
