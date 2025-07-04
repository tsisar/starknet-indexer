// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tsisar/starknet-indexer/generated/ent/protocolstat"
)

// ProtocolStatCreate is the builder for creating a ProtocolStat entity.
type ProtocolStatCreate struct {
	config
	mutation *ProtocolStatMutation
	hooks    []Hook
}

// SetTotalSupply sets the "totalSupply" field.
func (psc *ProtocolStatCreate) SetTotalSupply(s string) *ProtocolStatCreate {
	psc.mutation.SetTotalSupply(s)
	return psc
}

// SetTvl sets the "tvl" field.
func (psc *ProtocolStatCreate) SetTvl(s string) *ProtocolStatCreate {
	psc.mutation.SetTvl(s)
	return psc
}

// SetPools sets the "pools" field.
func (psc *ProtocolStatCreate) SetPools(s string) *ProtocolStatCreate {
	psc.mutation.SetPools(s)
	return psc
}

// SetID sets the "id" field.
func (psc *ProtocolStatCreate) SetID(s string) *ProtocolStatCreate {
	psc.mutation.SetID(s)
	return psc
}

// Mutation returns the ProtocolStatMutation object of the builder.
func (psc *ProtocolStatCreate) Mutation() *ProtocolStatMutation {
	return psc.mutation
}

// Save creates the ProtocolStat in the database.
func (psc *ProtocolStatCreate) Save(ctx context.Context) (*ProtocolStat, error) {
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProtocolStatCreate) SaveX(ctx context.Context) *ProtocolStat {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *ProtocolStatCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *ProtocolStatCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *ProtocolStatCreate) check() error {
	if _, ok := psc.mutation.TotalSupply(); !ok {
		return &ValidationError{Name: "totalSupply", err: errors.New(`ent: missing required field "ProtocolStat.totalSupply"`)}
	}
	if _, ok := psc.mutation.Tvl(); !ok {
		return &ValidationError{Name: "tvl", err: errors.New(`ent: missing required field "ProtocolStat.tvl"`)}
	}
	if _, ok := psc.mutation.Pools(); !ok {
		return &ValidationError{Name: "pools", err: errors.New(`ent: missing required field "ProtocolStat.pools"`)}
	}
	return nil
}

func (psc *ProtocolStatCreate) sqlSave(ctx context.Context) (*ProtocolStat, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ProtocolStat.ID type: %T", _spec.ID.Value)
		}
	}
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *ProtocolStatCreate) createSpec() (*ProtocolStat, *sqlgraph.CreateSpec) {
	var (
		_node = &ProtocolStat{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(protocolstat.Table, sqlgraph.NewFieldSpec(protocolstat.FieldID, field.TypeString))
	)
	if id, ok := psc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := psc.mutation.TotalSupply(); ok {
		_spec.SetField(protocolstat.FieldTotalSupply, field.TypeString, value)
		_node.TotalSupply = value
	}
	if value, ok := psc.mutation.Tvl(); ok {
		_spec.SetField(protocolstat.FieldTvl, field.TypeString, value)
		_node.Tvl = value
	}
	if value, ok := psc.mutation.Pools(); ok {
		_spec.SetField(protocolstat.FieldPools, field.TypeString, value)
		_node.Pools = value
	}
	return _node, _spec
}

// ProtocolStatCreateBulk is the builder for creating many ProtocolStat entities in bulk.
type ProtocolStatCreateBulk struct {
	config
	err      error
	builders []*ProtocolStatCreate
}

// Save creates the ProtocolStat entities in the database.
func (pscb *ProtocolStatCreateBulk) Save(ctx context.Context) ([]*ProtocolStat, error) {
	if pscb.err != nil {
		return nil, pscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*ProtocolStat, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProtocolStatMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *ProtocolStatCreateBulk) SaveX(ctx context.Context) []*ProtocolStat {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *ProtocolStatCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *ProtocolStatCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
