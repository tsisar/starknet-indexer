package schema

import "entgo.io/ent/schema/field"

var PoolOrderFieldEnum = field.Enum("pool_order_field").
	Values("id", "poolName", "debtCeiling", "liquidationRatio", "stabilityFeeRate", "tokenAdapterAddress", "lockedCollateral", "collateralPrice", "collateralLastPrice", "priceWithSafetyMargin", "rawPrice", "debtAccumulatedRate", "totalBorrowed", "totalAvailable", "tvl")
