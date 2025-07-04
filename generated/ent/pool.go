// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tsisar/starknet-indexer/generated/ent/pool"
)

// Pool is the model entity for the Pool schema.
type Pool struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// PoolName holds the value of the "poolName" field.
	PoolName string `json:"poolName,omitempty"`
	// DebtCeiling holds the value of the "debtCeiling" field.
	DebtCeiling string `json:"debtCeiling,omitempty"`
	// LiquidationRatio holds the value of the "liquidationRatio" field.
	LiquidationRatio string `json:"liquidationRatio,omitempty"`
	// StabilityFeeRate holds the value of the "stabilityFeeRate" field.
	StabilityFeeRate string `json:"stabilityFeeRate,omitempty"`
	// TokenAdapterAddress holds the value of the "tokenAdapterAddress" field.
	TokenAdapterAddress string `json:"tokenAdapterAddress,omitempty"`
	// LockedCollateral holds the value of the "lockedCollateral" field.
	LockedCollateral string `json:"lockedCollateral,omitempty"`
	// CollateralPrice holds the value of the "collateralPrice" field.
	CollateralPrice string `json:"collateralPrice,omitempty"`
	// CollateralLastPrice holds the value of the "collateralLastPrice" field.
	CollateralLastPrice string `json:"collateralLastPrice,omitempty"`
	// PriceWithSafetyMargin holds the value of the "priceWithSafetyMargin" field.
	PriceWithSafetyMargin string `json:"priceWithSafetyMargin,omitempty"`
	// RawPrice holds the value of the "rawPrice" field.
	RawPrice string `json:"rawPrice,omitempty"`
	// DebtAccumulatedRate holds the value of the "debtAccumulatedRate" field.
	DebtAccumulatedRate string `json:"debtAccumulatedRate,omitempty"`
	// TotalBorrowed holds the value of the "totalBorrowed" field.
	TotalBorrowed string `json:"totalBorrowed,omitempty"`
	// TotalAvailable holds the value of the "totalAvailable" field.
	TotalAvailable string `json:"totalAvailable,omitempty"`
	// Tvl holds the value of the "tvl" field.
	Tvl string `json:"tvl,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PoolQuery when eager-loading is set.
	Edges        PoolEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PoolEdges holds the relations/edges for other nodes in the graph.
type PoolEdges struct {
	// Positions holds the value of the positions edge.
	Positions []*Position `json:"positions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e PoolEdges) PositionsOrErr() ([]*Position, error) {
	if e.loadedTypes[0] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pool.FieldID, pool.FieldPoolName, pool.FieldDebtCeiling, pool.FieldLiquidationRatio, pool.FieldStabilityFeeRate, pool.FieldTokenAdapterAddress, pool.FieldLockedCollateral, pool.FieldCollateralPrice, pool.FieldCollateralLastPrice, pool.FieldPriceWithSafetyMargin, pool.FieldRawPrice, pool.FieldDebtAccumulatedRate, pool.FieldTotalBorrowed, pool.FieldTotalAvailable, pool.FieldTvl:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pool fields.
func (po *Pool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pool.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				po.ID = value.String
			}
		case pool.FieldPoolName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field poolName", values[i])
			} else if value.Valid {
				po.PoolName = value.String
			}
		case pool.FieldDebtCeiling:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field debtCeiling", values[i])
			} else if value.Valid {
				po.DebtCeiling = value.String
			}
		case pool.FieldLiquidationRatio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field liquidationRatio", values[i])
			} else if value.Valid {
				po.LiquidationRatio = value.String
			}
		case pool.FieldStabilityFeeRate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stabilityFeeRate", values[i])
			} else if value.Valid {
				po.StabilityFeeRate = value.String
			}
		case pool.FieldTokenAdapterAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tokenAdapterAddress", values[i])
			} else if value.Valid {
				po.TokenAdapterAddress = value.String
			}
		case pool.FieldLockedCollateral:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lockedCollateral", values[i])
			} else if value.Valid {
				po.LockedCollateral = value.String
			}
		case pool.FieldCollateralPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collateralPrice", values[i])
			} else if value.Valid {
				po.CollateralPrice = value.String
			}
		case pool.FieldCollateralLastPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collateralLastPrice", values[i])
			} else if value.Valid {
				po.CollateralLastPrice = value.String
			}
		case pool.FieldPriceWithSafetyMargin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field priceWithSafetyMargin", values[i])
			} else if value.Valid {
				po.PriceWithSafetyMargin = value.String
			}
		case pool.FieldRawPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rawPrice", values[i])
			} else if value.Valid {
				po.RawPrice = value.String
			}
		case pool.FieldDebtAccumulatedRate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field debtAccumulatedRate", values[i])
			} else if value.Valid {
				po.DebtAccumulatedRate = value.String
			}
		case pool.FieldTotalBorrowed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field totalBorrowed", values[i])
			} else if value.Valid {
				po.TotalBorrowed = value.String
			}
		case pool.FieldTotalAvailable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field totalAvailable", values[i])
			} else if value.Valid {
				po.TotalAvailable = value.String
			}
		case pool.FieldTvl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tvl", values[i])
			} else if value.Valid {
				po.Tvl = value.String
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pool.
// This includes values selected through modifiers, order, etc.
func (po *Pool) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryPositions queries the "positions" edge of the Pool entity.
func (po *Pool) QueryPositions() *PositionQuery {
	return NewPoolClient(po.config).QueryPositions(po)
}

// Update returns a builder for updating this Pool.
// Note that you need to call Pool.Unwrap() before calling this method if this Pool
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pool) Update() *PoolUpdateOne {
	return NewPoolClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Pool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pool) Unwrap() *Pool {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pool is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pool) String() string {
	var builder strings.Builder
	builder.WriteString("Pool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("poolName=")
	builder.WriteString(po.PoolName)
	builder.WriteString(", ")
	builder.WriteString("debtCeiling=")
	builder.WriteString(po.DebtCeiling)
	builder.WriteString(", ")
	builder.WriteString("liquidationRatio=")
	builder.WriteString(po.LiquidationRatio)
	builder.WriteString(", ")
	builder.WriteString("stabilityFeeRate=")
	builder.WriteString(po.StabilityFeeRate)
	builder.WriteString(", ")
	builder.WriteString("tokenAdapterAddress=")
	builder.WriteString(po.TokenAdapterAddress)
	builder.WriteString(", ")
	builder.WriteString("lockedCollateral=")
	builder.WriteString(po.LockedCollateral)
	builder.WriteString(", ")
	builder.WriteString("collateralPrice=")
	builder.WriteString(po.CollateralPrice)
	builder.WriteString(", ")
	builder.WriteString("collateralLastPrice=")
	builder.WriteString(po.CollateralLastPrice)
	builder.WriteString(", ")
	builder.WriteString("priceWithSafetyMargin=")
	builder.WriteString(po.PriceWithSafetyMargin)
	builder.WriteString(", ")
	builder.WriteString("rawPrice=")
	builder.WriteString(po.RawPrice)
	builder.WriteString(", ")
	builder.WriteString("debtAccumulatedRate=")
	builder.WriteString(po.DebtAccumulatedRate)
	builder.WriteString(", ")
	builder.WriteString("totalBorrowed=")
	builder.WriteString(po.TotalBorrowed)
	builder.WriteString(", ")
	builder.WriteString("totalAvailable=")
	builder.WriteString(po.TotalAvailable)
	builder.WriteString(", ")
	builder.WriteString("tvl=")
	builder.WriteString(po.Tvl)
	builder.WriteByte(')')
	return builder.String()
}

// Pools is a parsable slice of Pool.
type Pools []*Pool
