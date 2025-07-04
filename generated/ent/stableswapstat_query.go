// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tsisar/starknet-indexer/generated/ent/predicate"
	"github.com/tsisar/starknet-indexer/generated/ent/stableswapstat"
)

// StableSwapStatQuery is the builder for querying StableSwapStat entities.
type StableSwapStatQuery struct {
	config
	ctx        *QueryContext
	order      []stableswapstat.OrderOption
	inters     []Interceptor
	predicates []predicate.StableSwapStat
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StableSwapStatQuery builder.
func (sssq *StableSwapStatQuery) Where(ps ...predicate.StableSwapStat) *StableSwapStatQuery {
	sssq.predicates = append(sssq.predicates, ps...)
	return sssq
}

// Limit the number of records to be returned by this query.
func (sssq *StableSwapStatQuery) Limit(limit int) *StableSwapStatQuery {
	sssq.ctx.Limit = &limit
	return sssq
}

// Offset to start from.
func (sssq *StableSwapStatQuery) Offset(offset int) *StableSwapStatQuery {
	sssq.ctx.Offset = &offset
	return sssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sssq *StableSwapStatQuery) Unique(unique bool) *StableSwapStatQuery {
	sssq.ctx.Unique = &unique
	return sssq
}

// Order specifies how the records should be ordered.
func (sssq *StableSwapStatQuery) Order(o ...stableswapstat.OrderOption) *StableSwapStatQuery {
	sssq.order = append(sssq.order, o...)
	return sssq
}

// First returns the first StableSwapStat entity from the query.
// Returns a *NotFoundError when no StableSwapStat was found.
func (sssq *StableSwapStatQuery) First(ctx context.Context) (*StableSwapStat, error) {
	nodes, err := sssq.Limit(1).All(setContextOp(ctx, sssq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{stableswapstat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sssq *StableSwapStatQuery) FirstX(ctx context.Context) *StableSwapStat {
	node, err := sssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StableSwapStat ID from the query.
// Returns a *NotFoundError when no StableSwapStat ID was found.
func (sssq *StableSwapStatQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sssq.Limit(1).IDs(setContextOp(ctx, sssq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{stableswapstat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sssq *StableSwapStatQuery) FirstIDX(ctx context.Context) string {
	id, err := sssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StableSwapStat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StableSwapStat entity is found.
// Returns a *NotFoundError when no StableSwapStat entities are found.
func (sssq *StableSwapStatQuery) Only(ctx context.Context) (*StableSwapStat, error) {
	nodes, err := sssq.Limit(2).All(setContextOp(ctx, sssq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{stableswapstat.Label}
	default:
		return nil, &NotSingularError{stableswapstat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sssq *StableSwapStatQuery) OnlyX(ctx context.Context) *StableSwapStat {
	node, err := sssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StableSwapStat ID in the query.
// Returns a *NotSingularError when more than one StableSwapStat ID is found.
// Returns a *NotFoundError when no entities are found.
func (sssq *StableSwapStatQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sssq.Limit(2).IDs(setContextOp(ctx, sssq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{stableswapstat.Label}
	default:
		err = &NotSingularError{stableswapstat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sssq *StableSwapStatQuery) OnlyIDX(ctx context.Context) string {
	id, err := sssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StableSwapStats.
func (sssq *StableSwapStatQuery) All(ctx context.Context) ([]*StableSwapStat, error) {
	ctx = setContextOp(ctx, sssq.ctx, ent.OpQueryAll)
	if err := sssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StableSwapStat, *StableSwapStatQuery]()
	return withInterceptors[[]*StableSwapStat](ctx, sssq, qr, sssq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sssq *StableSwapStatQuery) AllX(ctx context.Context) []*StableSwapStat {
	nodes, err := sssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StableSwapStat IDs.
func (sssq *StableSwapStatQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sssq.ctx.Unique == nil && sssq.path != nil {
		sssq.Unique(true)
	}
	ctx = setContextOp(ctx, sssq.ctx, ent.OpQueryIDs)
	if err = sssq.Select(stableswapstat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sssq *StableSwapStatQuery) IDsX(ctx context.Context) []string {
	ids, err := sssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sssq *StableSwapStatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sssq.ctx, ent.OpQueryCount)
	if err := sssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sssq, querierCount[*StableSwapStatQuery](), sssq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sssq *StableSwapStatQuery) CountX(ctx context.Context) int {
	count, err := sssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sssq *StableSwapStatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sssq.ctx, ent.OpQueryExist)
	switch _, err := sssq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sssq *StableSwapStatQuery) ExistX(ctx context.Context) bool {
	exist, err := sssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StableSwapStatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sssq *StableSwapStatQuery) Clone() *StableSwapStatQuery {
	if sssq == nil {
		return nil
	}
	return &StableSwapStatQuery{
		config:     sssq.config,
		ctx:        sssq.ctx.Clone(),
		order:      append([]stableswapstat.OrderOption{}, sssq.order...),
		inters:     append([]Interceptor{}, sssq.inters...),
		predicates: append([]predicate.StableSwapStat{}, sssq.predicates...),
		// clone intermediate query.
		sql:  sssq.sql.Clone(),
		path: sssq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TotalTokenToStablecoinSwapEvents string `json:"totalTokenToStablecoinSwapEvents,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StableSwapStat.Query().
//		GroupBy(stableswapstat.FieldTotalTokenToStablecoinSwapEvents).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sssq *StableSwapStatQuery) GroupBy(field string, fields ...string) *StableSwapStatGroupBy {
	sssq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StableSwapStatGroupBy{build: sssq}
	grbuild.flds = &sssq.ctx.Fields
	grbuild.label = stableswapstat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TotalTokenToStablecoinSwapEvents string `json:"totalTokenToStablecoinSwapEvents,omitempty"`
//	}
//
//	client.StableSwapStat.Query().
//		Select(stableswapstat.FieldTotalTokenToStablecoinSwapEvents).
//		Scan(ctx, &v)
func (sssq *StableSwapStatQuery) Select(fields ...string) *StableSwapStatSelect {
	sssq.ctx.Fields = append(sssq.ctx.Fields, fields...)
	sbuild := &StableSwapStatSelect{StableSwapStatQuery: sssq}
	sbuild.label = stableswapstat.Label
	sbuild.flds, sbuild.scan = &sssq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StableSwapStatSelect configured with the given aggregations.
func (sssq *StableSwapStatQuery) Aggregate(fns ...AggregateFunc) *StableSwapStatSelect {
	return sssq.Select().Aggregate(fns...)
}

func (sssq *StableSwapStatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sssq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sssq); err != nil {
				return err
			}
		}
	}
	for _, f := range sssq.ctx.Fields {
		if !stableswapstat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sssq.path != nil {
		prev, err := sssq.path(ctx)
		if err != nil {
			return err
		}
		sssq.sql = prev
	}
	return nil
}

func (sssq *StableSwapStatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StableSwapStat, error) {
	var (
		nodes = []*StableSwapStat{}
		_spec = sssq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StableSwapStat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StableSwapStat{config: sssq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sssq *StableSwapStatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sssq.querySpec()
	_spec.Node.Columns = sssq.ctx.Fields
	if len(sssq.ctx.Fields) > 0 {
		_spec.Unique = sssq.ctx.Unique != nil && *sssq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sssq.driver, _spec)
}

func (sssq *StableSwapStatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(stableswapstat.Table, stableswapstat.Columns, sqlgraph.NewFieldSpec(stableswapstat.FieldID, field.TypeString))
	_spec.From = sssq.sql
	if unique := sssq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sssq.path != nil {
		_spec.Unique = true
	}
	if fields := sssq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stableswapstat.FieldID)
		for i := range fields {
			if fields[i] != stableswapstat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sssq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sssq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sssq *StableSwapStatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sssq.driver.Dialect())
	t1 := builder.Table(stableswapstat.Table)
	columns := sssq.ctx.Fields
	if len(columns) == 0 {
		columns = stableswapstat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sssq.sql != nil {
		selector = sssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sssq.ctx.Unique != nil && *sssq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sssq.predicates {
		p(selector)
	}
	for _, p := range sssq.order {
		p(selector)
	}
	if offset := sssq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sssq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StableSwapStatGroupBy is the group-by builder for StableSwapStat entities.
type StableSwapStatGroupBy struct {
	selector
	build *StableSwapStatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sssgb *StableSwapStatGroupBy) Aggregate(fns ...AggregateFunc) *StableSwapStatGroupBy {
	sssgb.fns = append(sssgb.fns, fns...)
	return sssgb
}

// Scan applies the selector query and scans the result into the given value.
func (sssgb *StableSwapStatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sssgb.build.ctx, ent.OpQueryGroupBy)
	if err := sssgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StableSwapStatQuery, *StableSwapStatGroupBy](ctx, sssgb.build, sssgb, sssgb.build.inters, v)
}

func (sssgb *StableSwapStatGroupBy) sqlScan(ctx context.Context, root *StableSwapStatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sssgb.fns))
	for _, fn := range sssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sssgb.flds)+len(sssgb.fns))
		for _, f := range *sssgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sssgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sssgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StableSwapStatSelect is the builder for selecting fields of StableSwapStat entities.
type StableSwapStatSelect struct {
	*StableSwapStatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ssss *StableSwapStatSelect) Aggregate(fns ...AggregateFunc) *StableSwapStatSelect {
	ssss.fns = append(ssss.fns, fns...)
	return ssss
}

// Scan applies the selector query and scans the result into the given value.
func (ssss *StableSwapStatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ssss.ctx, ent.OpQuerySelect)
	if err := ssss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StableSwapStatQuery, *StableSwapStatSelect](ctx, ssss.StableSwapStatQuery, ssss, ssss.inters, v)
}

func (ssss *StableSwapStatSelect) sqlScan(ctx context.Context, root *StableSwapStatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ssss.fns))
	for _, fn := range ssss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ssss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ssss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
