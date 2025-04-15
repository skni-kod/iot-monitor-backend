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
	"github.com/skni-kod/iot-monitor-backend/services/sensor-service/ent/predicate"
	"github.com/skni-kod/iot-monitor-backend/services/sensor-service/ent/sensor"
	"github.com/skni-kod/iot-monitor-backend/services/sensor-service/ent/sensortype"
)

// SensorQuery is the builder for querying Sensor entities.
type SensorQuery struct {
	config
	ctx        *QueryContext
	order      []sensor.OrderOption
	inters     []Interceptor
	predicates []predicate.Sensor
	withType   *SensorTypeQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SensorQuery builder.
func (sq *SensorQuery) Where(ps ...predicate.Sensor) *SensorQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SensorQuery) Limit(limit int) *SensorQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SensorQuery) Offset(offset int) *SensorQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SensorQuery) Unique(unique bool) *SensorQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SensorQuery) Order(o ...sensor.OrderOption) *SensorQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryType chains the current query on the "type" edge.
func (sq *SensorQuery) QueryType() *SensorTypeQuery {
	query := (&SensorTypeClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sensor.Table, sensor.FieldID, selector),
			sqlgraph.To(sensortype.Table, sensortype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, sensor.TypeTable, sensor.TypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Sensor entity from the query.
// Returns a *NotFoundError when no Sensor was found.
func (sq *SensorQuery) First(ctx context.Context) (*Sensor, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sensor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SensorQuery) FirstX(ctx context.Context) *Sensor {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Sensor ID from the query.
// Returns a *NotFoundError when no Sensor ID was found.
func (sq *SensorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sensor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SensorQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Sensor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Sensor entity is found.
// Returns a *NotFoundError when no Sensor entities are found.
func (sq *SensorQuery) Only(ctx context.Context) (*Sensor, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sensor.Label}
	default:
		return nil, &NotSingularError{sensor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SensorQuery) OnlyX(ctx context.Context) *Sensor {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Sensor ID in the query.
// Returns a *NotSingularError when more than one Sensor ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SensorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sensor.Label}
	default:
		err = &NotSingularError{sensor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SensorQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sensors.
func (sq *SensorQuery) All(ctx context.Context) ([]*Sensor, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Sensor, *SensorQuery]()
	return withInterceptors[[]*Sensor](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SensorQuery) AllX(ctx context.Context) []*Sensor {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Sensor IDs.
func (sq *SensorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(sensor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SensorQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SensorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SensorQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SensorQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SensorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SensorQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SensorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SensorQuery) Clone() *SensorQuery {
	if sq == nil {
		return nil
	}
	return &SensorQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]sensor.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Sensor{}, sq.predicates...),
		withType:   sq.withType.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithType tells the query-builder to eager-load the nodes that are connected to
// the "type" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SensorQuery) WithType(opts ...func(*SensorTypeQuery)) *SensorQuery {
	query := (&SensorTypeClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withType = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Sensor.Query().
//		GroupBy(sensor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SensorQuery) GroupBy(field string, fields ...string) *SensorGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SensorGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = sensor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Sensor.Query().
//		Select(sensor.FieldName).
//		Scan(ctx, &v)
func (sq *SensorQuery) Select(fields ...string) *SensorSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SensorSelect{SensorQuery: sq}
	sbuild.label = sensor.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SensorSelect configured with the given aggregations.
func (sq *SensorQuery) Aggregate(fns ...AggregateFunc) *SensorSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SensorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !sensor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SensorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Sensor, error) {
	var (
		nodes       = []*Sensor{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withType != nil,
		}
	)
	if sq.withType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sensor.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Sensor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Sensor{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withType; query != nil {
		if err := sq.loadType(ctx, query, nodes, nil,
			func(n *Sensor, e *SensorType) { n.Edges.Type = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SensorQuery) loadType(ctx context.Context, query *SensorTypeQuery, nodes []*Sensor, init func(*Sensor), assign func(*Sensor, *SensorType)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Sensor)
	for i := range nodes {
		if nodes[i].sensor_type == nil {
			continue
		}
		fk := *nodes[i].sensor_type
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sensortype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sensor_type" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SensorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SensorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sensor.Table, sensor.Columns, sqlgraph.NewFieldSpec(sensor.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sensor.FieldID)
		for i := range fields {
			if fields[i] != sensor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SensorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(sensor.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = sensor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SensorGroupBy is the group-by builder for Sensor entities.
type SensorGroupBy struct {
	selector
	build *SensorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SensorGroupBy) Aggregate(fns ...AggregateFunc) *SensorGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SensorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SensorQuery, *SensorGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SensorGroupBy) sqlScan(ctx context.Context, root *SensorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SensorSelect is the builder for selecting fields of Sensor entities.
type SensorSelect struct {
	*SensorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SensorSelect) Aggregate(fns ...AggregateFunc) *SensorSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SensorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SensorQuery, *SensorSelect](ctx, ss.SensorQuery, ss, ss.inters, v)
}

func (ss *SensorSelect) sqlScan(ctx context.Context, root *SensorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
