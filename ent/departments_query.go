// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"pamukkale_university/ent/departments"
	"pamukkale_university/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DepartmentsQuery is the builder for querying Departments entities.
type DepartmentsQuery struct {
	config
	ctx        *QueryContext
	order      []departments.Order
	inters     []Interceptor
	predicates []predicate.Departments
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DepartmentsQuery builder.
func (dq *DepartmentsQuery) Where(ps ...predicate.Departments) *DepartmentsQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DepartmentsQuery) Limit(limit int) *DepartmentsQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DepartmentsQuery) Offset(offset int) *DepartmentsQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DepartmentsQuery) Unique(unique bool) *DepartmentsQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DepartmentsQuery) Order(o ...departments.Order) *DepartmentsQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Departments entity from the query.
// Returns a *NotFoundError when no Departments was found.
func (dq *DepartmentsQuery) First(ctx context.Context) (*Departments, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{departments.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DepartmentsQuery) FirstX(ctx context.Context) *Departments {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Departments ID from the query.
// Returns a *NotFoundError when no Departments ID was found.
func (dq *DepartmentsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{departments.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DepartmentsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Departments entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Departments entity is found.
// Returns a *NotFoundError when no Departments entities are found.
func (dq *DepartmentsQuery) Only(ctx context.Context) (*Departments, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{departments.Label}
	default:
		return nil, &NotSingularError{departments.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DepartmentsQuery) OnlyX(ctx context.Context) *Departments {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Departments ID in the query.
// Returns a *NotSingularError when more than one Departments ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DepartmentsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{departments.Label}
	default:
		err = &NotSingularError{departments.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DepartmentsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DepartmentsSlice.
func (dq *DepartmentsQuery) All(ctx context.Context) ([]*Departments, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Departments, *DepartmentsQuery]()
	return withInterceptors[[]*Departments](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DepartmentsQuery) AllX(ctx context.Context) []*Departments {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Departments IDs.
func (dq *DepartmentsQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(departments.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DepartmentsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DepartmentsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DepartmentsQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DepartmentsQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DepartmentsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DepartmentsQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DepartmentsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DepartmentsQuery) Clone() *DepartmentsQuery {
	if dq == nil {
		return nil
	}
	return &DepartmentsQuery{
		config:     dq.config,
		ctx:        dq.ctx.Clone(),
		order:      append([]departments.Order{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Departments{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Departments.Query().
//		GroupBy(departments.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DepartmentsQuery) GroupBy(field string, fields ...string) *DepartmentsGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DepartmentsGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = departments.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//	}
//
//	client.Departments.Query().
//		Select(departments.FieldCode).
//		Scan(ctx, &v)
func (dq *DepartmentsQuery) Select(fields ...string) *DepartmentsSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DepartmentsSelect{DepartmentsQuery: dq}
	sbuild.label = departments.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DepartmentsSelect configured with the given aggregations.
func (dq *DepartmentsQuery) Aggregate(fns ...AggregateFunc) *DepartmentsSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DepartmentsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !departments.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DepartmentsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Departments, error) {
	var (
		nodes = []*Departments{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Departments).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Departments{config: dq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dq *DepartmentsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DepartmentsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(departments.Table, departments.Columns, sqlgraph.NewFieldSpec(departments.FieldID, field.TypeUUID))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, departments.FieldID)
		for i := range fields {
			if fields[i] != departments.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DepartmentsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(departments.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = departments.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DepartmentsGroupBy is the group-by builder for Departments entities.
type DepartmentsGroupBy struct {
	selector
	build *DepartmentsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DepartmentsGroupBy) Aggregate(fns ...AggregateFunc) *DepartmentsGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DepartmentsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentsQuery, *DepartmentsGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DepartmentsGroupBy) sqlScan(ctx context.Context, root *DepartmentsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DepartmentsSelect is the builder for selecting fields of Departments entities.
type DepartmentsSelect struct {
	*DepartmentsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DepartmentsSelect) Aggregate(fns ...AggregateFunc) *DepartmentsSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DepartmentsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DepartmentsQuery, *DepartmentsSelect](ctx, ds.DepartmentsQuery, ds, ds.inters, v)
}

func (ds *DepartmentsSelect) sqlScan(ctx context.Context, root *DepartmentsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}