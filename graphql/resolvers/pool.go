package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/pool"
	"github.com/tsisar/starknet-indexer/graphql/model"
)

func ApplyPoolWhereInput(query *ent.PoolQuery, where *model.PoolWhereInput) *ent.PoolQuery {
	if where == nil {
		return query
	}

	if where.ID != nil {
		query = query.Where(pool.IDEQ(*where.ID))
	}
	if where.IDNot != nil {
		query = query.Where(pool.IDNEQ(*where.IDNot))
	}
	if where.IDIn != nil {
		query = query.Where(pool.IDIn(where.IDIn...))
	}
	if where.IDNotIn != nil {
		query = query.Where(pool.IDNotIn(where.IDNotIn...))
	}
	if where.PoolName != nil {
		query = query.Where(pool.PoolNameEQ(*where.PoolName))
	}
	if where.PoolNameNot != nil {
		query = query.Where(pool.PoolNameNEQ(*where.PoolNameNot))
	}
	if where.PoolNameIn != nil {
		query = query.Where(pool.PoolNameIn(where.PoolNameIn...))
	}
	if where.PoolNameNotIn != nil {
		query = query.Where(pool.PoolNameNotIn(where.PoolNameNotIn...))
	}
	if where.DebtCeiling != nil {
		query = query.Where(pool.DebtCeilingEQ(*where.DebtCeiling))
	}
	if where.DebtCeilingNot != nil {
		query = query.Where(pool.DebtCeilingNEQ(*where.DebtCeilingNot))
	}
	if where.DebtCeilingIn != nil {
		query = query.Where(pool.DebtCeilingIn(where.DebtCeilingIn...))
	}
	if where.DebtCeilingNotIn != nil {
		query = query.Where(pool.DebtCeilingNotIn(where.DebtCeilingNotIn...))
	}
	if where.LiquidationRatio != nil {
		query = query.Where(pool.LiquidationRatioEQ(*where.LiquidationRatio))
	}
	if where.LiquidationRatioNot != nil {
		query = query.Where(pool.LiquidationRatioNEQ(*where.LiquidationRatioNot))
	}
	if where.LiquidationRatioIn != nil {
		query = query.Where(pool.LiquidationRatioIn(where.LiquidationRatioIn...))
	}
	if where.LiquidationRatioNotIn != nil {
		query = query.Where(pool.LiquidationRatioNotIn(where.LiquidationRatioNotIn...))
	}
	if where.StabilityFeeRate != nil {
		query = query.Where(pool.StabilityFeeRateEQ(*where.StabilityFeeRate))
	}
	if where.StabilityFeeRateNot != nil {
		query = query.Where(pool.StabilityFeeRateNEQ(*where.StabilityFeeRateNot))
	}
	if where.StabilityFeeRateIn != nil {
		query = query.Where(pool.StabilityFeeRateIn(where.StabilityFeeRateIn...))
	}
	if where.StabilityFeeRateNotIn != nil {
		query = query.Where(pool.StabilityFeeRateNotIn(where.StabilityFeeRateNotIn...))
	}
	if where.TokenAdapterAddress != nil {
		query = query.Where(pool.TokenAdapterAddressEQ(*where.TokenAdapterAddress))
	}
	if where.TokenAdapterAddressNot != nil {
		query = query.Where(pool.TokenAdapterAddressNEQ(*where.TokenAdapterAddressNot))
	}
	if where.TokenAdapterAddressIn != nil {
		query = query.Where(pool.TokenAdapterAddressIn(where.TokenAdapterAddressIn...))
	}
	if where.TokenAdapterAddressNotIn != nil {
		query = query.Where(pool.TokenAdapterAddressNotIn(where.TokenAdapterAddressNotIn...))
	}
	if where.LockedCollateral != nil {
		query = query.Where(pool.LockedCollateralEQ(*where.LockedCollateral))
	}
	if where.LockedCollateralNot != nil {
		query = query.Where(pool.LockedCollateralNEQ(*where.LockedCollateralNot))
	}
	if where.LockedCollateralIn != nil {
		query = query.Where(pool.LockedCollateralIn(where.LockedCollateralIn...))
	}
	if where.LockedCollateralNotIn != nil {
		query = query.Where(pool.LockedCollateralNotIn(where.LockedCollateralNotIn...))
	}
	if where.CollateralPrice != nil {
		query = query.Where(pool.CollateralPriceEQ(*where.CollateralPrice))
	}
	if where.CollateralPriceNot != nil {
		query = query.Where(pool.CollateralPriceNEQ(*where.CollateralPriceNot))
	}
	if where.CollateralPriceIn != nil {
		query = query.Where(pool.CollateralPriceIn(where.CollateralPriceIn...))
	}
	if where.CollateralPriceNotIn != nil {
		query = query.Where(pool.CollateralPriceNotIn(where.CollateralPriceNotIn...))
	}
	if where.CollateralLastPrice != nil {
		query = query.Where(pool.CollateralLastPriceEQ(*where.CollateralLastPrice))
	}
	if where.CollateralLastPriceNot != nil {
		query = query.Where(pool.CollateralLastPriceNEQ(*where.CollateralLastPriceNot))
	}
	if where.CollateralLastPriceIn != nil {
		query = query.Where(pool.CollateralLastPriceIn(where.CollateralLastPriceIn...))
	}
	if where.CollateralLastPriceNotIn != nil {
		query = query.Where(pool.CollateralLastPriceNotIn(where.CollateralLastPriceNotIn...))
	}
	if where.PriceWithSafetyMargin != nil {
		query = query.Where(pool.PriceWithSafetyMarginEQ(*where.PriceWithSafetyMargin))
	}
	if where.PriceWithSafetyMarginNot != nil {
		query = query.Where(pool.PriceWithSafetyMarginNEQ(*where.PriceWithSafetyMarginNot))
	}
	if where.PriceWithSafetyMarginIn != nil {
		query = query.Where(pool.PriceWithSafetyMarginIn(where.PriceWithSafetyMarginIn...))
	}
	if where.PriceWithSafetyMarginNotIn != nil {
		query = query.Where(pool.PriceWithSafetyMarginNotIn(where.PriceWithSafetyMarginNotIn...))
	}
	if where.RawPrice != nil {
		query = query.Where(pool.RawPriceEQ(*where.RawPrice))
	}
	if where.RawPriceNot != nil {
		query = query.Where(pool.RawPriceNEQ(*where.RawPriceNot))
	}
	if where.RawPriceIn != nil {
		query = query.Where(pool.RawPriceIn(where.RawPriceIn...))
	}
	if where.RawPriceNotIn != nil {
		query = query.Where(pool.RawPriceNotIn(where.RawPriceNotIn...))
	}
	if where.DebtAccumulatedRate != nil {
		query = query.Where(pool.DebtAccumulatedRateEQ(*where.DebtAccumulatedRate))
	}
	if where.DebtAccumulatedRateNot != nil {
		query = query.Where(pool.DebtAccumulatedRateNEQ(*where.DebtAccumulatedRateNot))
	}
	if where.DebtAccumulatedRateIn != nil {
		query = query.Where(pool.DebtAccumulatedRateIn(where.DebtAccumulatedRateIn...))
	}
	if where.DebtAccumulatedRateNotIn != nil {
		query = query.Where(pool.DebtAccumulatedRateNotIn(where.DebtAccumulatedRateNotIn...))
	}
	if where.TotalBorrowed != nil {
		query = query.Where(pool.TotalBorrowedEQ(*where.TotalBorrowed))
	}
	if where.TotalBorrowedNot != nil {
		query = query.Where(pool.TotalBorrowedNEQ(*where.TotalBorrowedNot))
	}
	if where.TotalBorrowedIn != nil {
		query = query.Where(pool.TotalBorrowedIn(where.TotalBorrowedIn...))
	}
	if where.TotalBorrowedNotIn != nil {
		query = query.Where(pool.TotalBorrowedNotIn(where.TotalBorrowedNotIn...))
	}
	if where.TotalAvailable != nil {
		query = query.Where(pool.TotalAvailableEQ(*where.TotalAvailable))
	}
	if where.TotalAvailableNot != nil {
		query = query.Where(pool.TotalAvailableNEQ(*where.TotalAvailableNot))
	}
	if where.TotalAvailableIn != nil {
		query = query.Where(pool.TotalAvailableIn(where.TotalAvailableIn...))
	}
	if where.TotalAvailableNotIn != nil {
		query = query.Where(pool.TotalAvailableNotIn(where.TotalAvailableNotIn...))
	}
	if where.Tvl != nil {
		query = query.Where(pool.TvlEQ(*where.Tvl))
	}
	if where.TvlNot != nil {
		query = query.Where(pool.TvlNEQ(*where.TvlNot))
	}
	if where.TvlIn != nil {
		query = query.Where(pool.TvlIn(where.TvlIn...))
	}
	if where.TvlNotIn != nil {
		query = query.Where(pool.TvlNotIn(where.TvlNotIn...))
	}

	return query
}

func ApplyPoolOrderBy(query *ent.PoolQuery, orderBy *model.PoolOrderBy) *ent.PoolQuery {
	if orderBy == nil {
		return query
	}

	field := orderBy.Field
	direction := orderBy.Direction

	orderFunc := ent.Asc
	if direction == model.OrderDirectionDesc {
		orderFunc = ent.Desc
	}

	switch field {
	case model.PoolOrderFieldID:
		query = query.Order(orderFunc(pool.FieldID))
	case model.PoolOrderFieldPoolName:
		query = query.Order(orderFunc(pool.FieldPoolName))
	case model.PoolOrderFieldDebtCeiling:
		query = query.Order(orderFunc(pool.FieldDebtCeiling))
	case model.PoolOrderFieldLiquidationRatio:
		query = query.Order(orderFunc(pool.FieldLiquidationRatio))
	case model.PoolOrderFieldStabilityFeeRate:
		query = query.Order(orderFunc(pool.FieldStabilityFeeRate))
	case model.PoolOrderFieldTokenAdapterAddress:
		query = query.Order(orderFunc(pool.FieldTokenAdapterAddress))
	case model.PoolOrderFieldLockedCollateral:
		query = query.Order(orderFunc(pool.FieldLockedCollateral))
	case model.PoolOrderFieldCollateralPrice:
		query = query.Order(orderFunc(pool.FieldCollateralPrice))
	case model.PoolOrderFieldCollateralLastPrice:
		query = query.Order(orderFunc(pool.FieldCollateralLastPrice))
	case model.PoolOrderFieldPriceWithSafetyMargin:
		query = query.Order(orderFunc(pool.FieldPriceWithSafetyMargin))
	case model.PoolOrderFieldRawPrice:
		query = query.Order(orderFunc(pool.FieldRawPrice))
	case model.PoolOrderFieldDebtAccumulatedRate:
		query = query.Order(orderFunc(pool.FieldDebtAccumulatedRate))
	case model.PoolOrderFieldTotalBorrowed:
		query = query.Order(orderFunc(pool.FieldTotalBorrowed))
	case model.PoolOrderFieldTotalAvailable:
		query = query.Order(orderFunc(pool.FieldTotalAvailable))
	case model.PoolOrderFieldTvl:
		query = query.Order(orderFunc(pool.FieldTvl))
	}

	return query
}

func ApplyPoolLimit(query *ent.PoolQuery, first, skip *int32) *ent.PoolQuery {
	if skip != nil {
		query = query.Offset(int(*skip))
	}
	if first != nil {
		query = query.Limit(int(*first))
	} else {
		query = query.Limit(100)
	}
	return query
}
