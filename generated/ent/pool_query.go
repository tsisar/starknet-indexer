// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tsisar/starknet-indexer/generated/ent/pool"
	"github.com/tsisar/starknet-indexer/generated/ent/position"
	"github.com/tsisar/starknet-indexer/generated/ent/predicate"
)

// PoolQuery is the builder for querying Pool entities.
type PoolQuery struct {
	config
	ctx           *QueryContext
	order         []pool.OrderOption
	inters        []Interceptor
	predicates    []predicate.Pool
	withPositions *PositionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PoolQuery builder.
func (pq *PoolQuery) Where(ps ...predicate.Pool) *PoolQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PoolQuery) Limit(limit int) *PoolQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PoolQuery) Offset(offset int) *PoolQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PoolQuery) Unique(unique bool) *PoolQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PoolQuery) Order(o ...pool.OrderOption) *PoolQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPositions chains the current query on the "positions" edge.
func (pq *PoolQuery) QueryPositions() *PositionQuery {
	query := (&PositionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pool.Table, pool.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pool.PositionsTable, pool.PositionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pool entity from the query.
// Returns a *NotFoundError when no Pool was found.
func (pq *PoolQuery) First(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PoolQuery) FirstX(ctx context.Context) *Pool {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pool ID from the query.
// Returns a *NotFoundError when no Pool ID was found.
func (pq *PoolQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PoolQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pool entity is found.
// Returns a *NotFoundError when no Pool entities are found.
func (pq *PoolQuery) Only(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pool.Label}
	default:
		return nil, &NotSingularError{pool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PoolQuery) OnlyX(ctx context.Context) *Pool {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pool ID in the query.
// Returns a *NotSingularError when more than one Pool ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PoolQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pool.Label}
	default:
		err = &NotSingularError{pool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PoolQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pools.
func (pq *PoolQuery) All(ctx context.Context) ([]*Pool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pool, *PoolQuery]()
	return withInterceptors[[]*Pool](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PoolQuery) AllX(ctx context.Context) []*Pool {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pool IDs.
func (pq *PoolQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(pool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PoolQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PoolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PoolQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PoolQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PoolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PoolQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PoolQuery) Clone() *PoolQuery {
	if pq == nil {
		return nil
	}
	return &PoolQuery{
		config:        pq.config,
		ctx:           pq.ctx.Clone(),
		order:         append([]pool.OrderOption{}, pq.order...),
		inters:        append([]Interceptor{}, pq.inters...),
		predicates:    append([]predicate.Pool{}, pq.predicates...),
		withPositions: pq.withPositions.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPositions tells the query-builder to eager-load the nodes that are connected to
// the "positions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithPositions(opts ...func(*PositionQuery)) *PoolQuery {
	query := (&PositionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPositions = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PoolName string `json:"poolName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pool.Query().
//		GroupBy(pool.FieldPoolName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PoolQuery) GroupBy(field string, fields ...string) *PoolGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PoolGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pool.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PoolName string `json:"poolName,omitempty"`
//	}
//
//	client.Pool.Query().
//		Select(pool.FieldPoolName).
//		Scan(ctx, &v)
func (pq *PoolQuery) Select(fields ...string) *PoolSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PoolSelect{PoolQuery: pq}
	sbuild.label = pool.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PoolSelect configured with the given aggregations.
func (pq *PoolQuery) Aggregate(fns ...AggregateFunc) *PoolSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PoolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pool, error) {
	var (
		nodes       = []*Pool{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withPositions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pool{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPositions; query != nil {
		if err := pq.loadPositions(ctx, query, nodes,
			func(n *Pool) { n.Edges.Positions = []*Position{} },
			func(n *Pool, e *Position) { n.Edges.Positions = append(n.Edges.Positions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PoolQuery) loadPositions(ctx context.Context, query *PositionQuery, nodes []*Pool, init func(*Pool), assign func(*Pool, *Position)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Pool)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Position(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(pool.PositionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.pool_positions
		if fk == nil {
			return fmt.Errorf(`foreign-key "pool_positions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "pool_positions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pool.Table, pool.Columns, sqlgraph.NewFieldSpec(pool.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for i := range fields {
			if fields[i] != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pool.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PoolGroupBy is the group-by builder for Pool entities.
type PoolGroupBy struct {
	selector
	build *PoolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PoolGroupBy) Aggregate(fns ...AggregateFunc) *PoolGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PoolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoolQuery, *PoolGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PoolGroupBy) sqlScan(ctx context.Context, root *PoolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PoolSelect is the builder for selecting fields of Pool entities.
type PoolSelect struct {
	*PoolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PoolSelect) Aggregate(fns ...AggregateFunc) *PoolSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PoolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoolQuery, *PoolSelect](ctx, ps.PoolQuery, ps, ps.inters, v)
}

func (ps *PoolSelect) sqlScan(ctx context.Context, root *PoolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
