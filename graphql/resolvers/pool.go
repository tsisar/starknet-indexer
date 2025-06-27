package resolvers

import (
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/generated/ent/pool"
	"github.com/tsisar/starknet-indexer/generated/ent/position"
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

	// Handle nested positions filters
	if where.Positions != nil {
		if where.Positions.ID != nil {
			query = query.Where(pool.HasPositionsWith(position.IDEQ(*where.Positions.ID)))
		}
		if where.Positions.IDNot != nil {
			query = query.Where(pool.HasPositionsWith(position.IDNEQ(*where.Positions.IDNot)))
		}
		if where.Positions.IDIn != nil {
			query = query.Where(pool.HasPositionsWith(position.IDIn(where.Positions.IDIn...)))
		}
		if where.Positions.IDNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.IDNotIn(where.Positions.IDNotIn...)))
		}
		if where.Positions.PositionAddress != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionAddressEQ(*where.Positions.PositionAddress)))
		}
		if where.Positions.PositionAddressNot != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionAddressNEQ(*where.Positions.PositionAddressNot)))
		}
		if where.Positions.PositionAddressIn != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionAddressIn(where.Positions.PositionAddressIn...)))
		}
		if where.Positions.PositionAddressNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionAddressNotIn(where.Positions.PositionAddressNotIn...)))
		}
		if where.Positions.UserAddress != nil {
			query = query.Where(pool.HasPositionsWith(position.UserAddressEQ(*where.Positions.UserAddress)))
		}
		if where.Positions.UserAddressNot != nil {
			query = query.Where(pool.HasPositionsWith(position.UserAddressNEQ(*where.Positions.UserAddressNot)))
		}
		if where.Positions.UserAddressIn != nil {
			query = query.Where(pool.HasPositionsWith(position.UserAddressIn(where.Positions.UserAddressIn...)))
		}
		if where.Positions.UserAddressNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.UserAddressNotIn(where.Positions.UserAddressNotIn...)))
		}
		if where.Positions.WalletAddress != nil {
			query = query.Where(pool.HasPositionsWith(position.WalletAddressEQ(*where.Positions.WalletAddress)))
		}
		if where.Positions.WalletAddressNot != nil {
			query = query.Where(pool.HasPositionsWith(position.WalletAddressNEQ(*where.Positions.WalletAddressNot)))
		}
		if where.Positions.WalletAddressIn != nil {
			query = query.Where(pool.HasPositionsWith(position.WalletAddressIn(where.Positions.WalletAddressIn...)))
		}
		if where.Positions.WalletAddressNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.WalletAddressNotIn(where.Positions.WalletAddressNotIn...)))
		}
		if where.Positions.CollateralPool != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolEQ(*where.Positions.CollateralPool)))
		}
		if where.Positions.CollateralPoolNot != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNEQ(*where.Positions.CollateralPoolNot)))
		}
		if where.Positions.CollateralPoolIn != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolIn(where.Positions.CollateralPoolIn...)))
		}
		if where.Positions.CollateralPoolNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNotIn(where.Positions.CollateralPoolNotIn...)))
		}
		if where.Positions.CollateralPoolName != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNameEQ(*where.Positions.CollateralPoolName)))
		}
		if where.Positions.CollateralPoolNameNot != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNameNEQ(*where.Positions.CollateralPoolNameNot)))
		}
		if where.Positions.CollateralPoolNameIn != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNameIn(where.Positions.CollateralPoolNameIn...)))
		}
		if where.Positions.CollateralPoolNameNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.CollateralPoolNameNotIn(where.Positions.CollateralPoolNameNotIn...)))
		}
		if where.Positions.PositionID != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionIdEQ(*where.Positions.PositionID)))
		}
		if where.Positions.PositionIDNot != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionIdNEQ(*where.Positions.PositionIDNot)))
		}
		if where.Positions.PositionIDIn != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionIdIn(where.Positions.PositionIDIn...)))
		}
		if where.Positions.PositionIDNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionIdNotIn(where.Positions.PositionIDNotIn...)))
		}
		if where.Positions.LockedCollateral != nil {
			query = query.Where(pool.HasPositionsWith(position.LockedCollateralEQ(*where.Positions.LockedCollateral)))
		}
		if where.Positions.LockedCollateralNot != nil {
			query = query.Where(pool.HasPositionsWith(position.LockedCollateralNEQ(*where.Positions.LockedCollateralNot)))
		}
		if where.Positions.LockedCollateralIn != nil {
			query = query.Where(pool.HasPositionsWith(position.LockedCollateralIn(where.Positions.LockedCollateralIn...)))
		}
		if where.Positions.LockedCollateralNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.LockedCollateralNotIn(where.Positions.LockedCollateralNotIn...)))
		}
		if where.Positions.DebtValue != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtValueEQ(*where.Positions.DebtValue)))
		}
		if where.Positions.DebtValueNot != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtValueNEQ(*where.Positions.DebtValueNot)))
		}
		if where.Positions.DebtValueIn != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtValueIn(where.Positions.DebtValueIn...)))
		}
		if where.Positions.DebtValueNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtValueNotIn(where.Positions.DebtValueNotIn...)))
		}
		if where.Positions.DebtShare != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtShareEQ(*where.Positions.DebtShare)))
		}
		if where.Positions.DebtShareNot != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtShareNEQ(*where.Positions.DebtShareNot)))
		}
		if where.Positions.DebtShareIn != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtShareIn(where.Positions.DebtShareIn...)))
		}
		if where.Positions.DebtShareNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.DebtShareNotIn(where.Positions.DebtShareNotIn...)))
		}
		if where.Positions.SafetyBuffer != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferEQ(*where.Positions.SafetyBuffer)))
		}
		if where.Positions.SafetyBufferNot != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferNEQ(*where.Positions.SafetyBufferNot)))
		}
		if where.Positions.SafetyBufferIn != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferIn(where.Positions.SafetyBufferIn...)))
		}
		if where.Positions.SafetyBufferNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferNotIn(where.Positions.SafetyBufferNotIn...)))
		}
		if where.Positions.SafetyBufferInPercent != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferInPercentEQ(*where.Positions.SafetyBufferInPercent)))
		}
		if where.Positions.SafetyBufferInPercentNot != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferInPercentNEQ(*where.Positions.SafetyBufferInPercentNot)))
		}
		if where.Positions.SafetyBufferInPercentIn != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferInPercentIn(where.Positions.SafetyBufferInPercentIn...)))
		}
		if where.Positions.SafetyBufferInPercentNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.SafetyBufferInPercentNotIn(where.Positions.SafetyBufferInPercentNotIn...)))
		}
		if where.Positions.Tvl != nil {
			query = query.Where(pool.HasPositionsWith(position.TvlEQ(*where.Positions.Tvl)))
		}
		if where.Positions.TvlNot != nil {
			query = query.Where(pool.HasPositionsWith(position.TvlNEQ(*where.Positions.TvlNot)))
		}
		if where.Positions.TvlIn != nil {
			query = query.Where(pool.HasPositionsWith(position.TvlIn(where.Positions.TvlIn...)))
		}
		if where.Positions.TvlNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.TvlNotIn(where.Positions.TvlNotIn...)))
		}
		if where.Positions.PositionStatus != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionStatusEQ(string(*where.Positions.PositionStatus))))
		}
		if where.Positions.PositionStatusNot != nil {
			query = query.Where(pool.HasPositionsWith(position.PositionStatusNEQ(string(*where.Positions.PositionStatusNot))))
		}
		if where.Positions.PositionStatusIn != nil {
			statuses := make([]string, len(where.Positions.PositionStatusIn))
			for i, status := range where.Positions.PositionStatusIn {
				statuses[i] = string(status)
			}
			query = query.Where(pool.HasPositionsWith(position.PositionStatusIn(statuses...)))
		}
		if where.Positions.PositionStatusNotIn != nil {
			statuses := make([]string, len(where.Positions.PositionStatusNotIn))
			for i, status := range where.Positions.PositionStatusNotIn {
				statuses[i] = string(status)
			}
			query = query.Where(pool.HasPositionsWith(position.PositionStatusNotIn(statuses...)))
		}
		if where.Positions.LiquidationCount != nil {
			query = query.Where(pool.HasPositionsWith(position.LiquidationCountEQ(*where.Positions.LiquidationCount)))
		}
		if where.Positions.LiquidationCountNot != nil {
			query = query.Where(pool.HasPositionsWith(position.LiquidationCountNEQ(*where.Positions.LiquidationCountNot)))
		}
		if where.Positions.LiquidationCountIn != nil {
			query = query.Where(pool.HasPositionsWith(position.LiquidationCountIn(where.Positions.LiquidationCountIn...)))
		}
		if where.Positions.LiquidationCountNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.LiquidationCountNotIn(where.Positions.LiquidationCountNotIn...)))
		}
		if where.Positions.BlockNumber != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockNumberEQ(*where.Positions.BlockNumber)))
		}
		if where.Positions.BlockNumberNot != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockNumberNEQ(*where.Positions.BlockNumberNot)))
		}
		if where.Positions.BlockNumberIn != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockNumberIn(where.Positions.BlockNumberIn...)))
		}
		if where.Positions.BlockNumberNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockNumberNotIn(where.Positions.BlockNumberNotIn...)))
		}
		if where.Positions.BlockTimestamp != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockTimestampEQ(*where.Positions.BlockTimestamp)))
		}
		if where.Positions.BlockTimestampNot != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockTimestampNEQ(*where.Positions.BlockTimestampNot)))
		}
		if where.Positions.BlockTimestampIn != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockTimestampIn(where.Positions.BlockTimestampIn...)))
		}
		if where.Positions.BlockTimestampNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.BlockTimestampNotIn(where.Positions.BlockTimestampNotIn...)))
		}
		if where.Positions.Transaction != nil {
			query = query.Where(pool.HasPositionsWith(position.TransactionEQ(*where.Positions.Transaction)))
		}
		if where.Positions.TransactionNot != nil {
			query = query.Where(pool.HasPositionsWith(position.TransactionNEQ(*where.Positions.TransactionNot)))
		}
		if where.Positions.TransactionIn != nil {
			query = query.Where(pool.HasPositionsWith(position.TransactionIn(where.Positions.TransactionIn...)))
		}
		if where.Positions.TransactionNotIn != nil {
			query = query.Where(pool.HasPositionsWith(position.TransactionNotIn(where.Positions.TransactionNotIn...)))
		}
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
