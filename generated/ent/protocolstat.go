// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tsisar/starknet-indexer/generated/ent/protocolstat"
)

// ProtocolStat is the model entity for the ProtocolStat schema.
type ProtocolStat struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TotalSupply holds the value of the "totalSupply" field.
	TotalSupply string `json:"totalSupply,omitempty"`
	// Tvl holds the value of the "tvl" field.
	Tvl string `json:"tvl,omitempty"`
	// Pools holds the value of the "pools" field.
	Pools        string `json:"pools,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProtocolStat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case protocolstat.FieldID, protocolstat.FieldTotalSupply, protocolstat.FieldTvl, protocolstat.FieldPools:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProtocolStat fields.
func (ps *ProtocolStat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case protocolstat.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ps.ID = value.String
			}
		case protocolstat.FieldTotalSupply:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field totalSupply", values[i])
			} else if value.Valid {
				ps.TotalSupply = value.String
			}
		case protocolstat.FieldTvl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tvl", values[i])
			} else if value.Valid {
				ps.Tvl = value.String
			}
		case protocolstat.FieldPools:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pools", values[i])
			} else if value.Valid {
				ps.Pools = value.String
			}
		default:
			ps.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProtocolStat.
// This includes values selected through modifiers, order, etc.
func (ps *ProtocolStat) Value(name string) (ent.Value, error) {
	return ps.selectValues.Get(name)
}

// Update returns a builder for updating this ProtocolStat.
// Note that you need to call ProtocolStat.Unwrap() before calling this method if this ProtocolStat
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *ProtocolStat) Update() *ProtocolStatUpdateOne {
	return NewProtocolStatClient(ps.config).UpdateOne(ps)
}

// Unwrap unwraps the ProtocolStat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *ProtocolStat) Unwrap() *ProtocolStat {
	_tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProtocolStat is not a transactional entity")
	}
	ps.config.driver = _tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *ProtocolStat) String() string {
	var builder strings.Builder
	builder.WriteString("ProtocolStat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ps.ID))
	builder.WriteString("totalSupply=")
	builder.WriteString(ps.TotalSupply)
	builder.WriteString(", ")
	builder.WriteString("tvl=")
	builder.WriteString(ps.Tvl)
	builder.WriteString(", ")
	builder.WriteString("pools=")
	builder.WriteString(ps.Pools)
	builder.WriteByte(')')
	return builder.String()
}

// ProtocolStats is a parsable slice of ProtocolStat.
type ProtocolStats []*ProtocolStat
