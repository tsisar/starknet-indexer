package schema

import "entgo.io/ent/schema/field"

var PositionOrderFieldEnum = field.Enum("position_order_field").
	Values("id", "positionAddress", "userAddress", "walletAddress", "collateralPool", "collateralPoolName", "positionId", "lockedCollateral", "debtValue", "debtShare", "safetyBuffer", "safetyBufferInPercent", "tvl", "positionStatus", "liquidationCount", "blockNumber", "blockTimestamp", "transaction", "pool")
