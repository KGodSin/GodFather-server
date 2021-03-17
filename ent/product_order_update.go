// Code generated by entc, DO NOT EDIT.

package ent

import (
	"GodFather-server/ent/predicate"
	"GodFather-server/ent/product_order"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductOrderUpdate is the builder for updating Product_order entities.
type ProductOrderUpdate struct {
	config
	hooks    []Hook
	mutation *ProductOrderMutation
}

// Where adds a new predicate for the ProductOrderUpdate builder.
func (pou *ProductOrderUpdate) Where(ps ...predicate.Product_order) *ProductOrderUpdate {
	pou.mutation.predicates = append(pou.mutation.predicates, ps...)
	return pou
}

// SetName sets the "name" field.
func (pou *ProductOrderUpdate) SetName(s string) *ProductOrderUpdate {
	pou.mutation.SetName(s)
	return pou
}

// SetPrice sets the "price" field.
func (pou *ProductOrderUpdate) SetPrice(i int) *ProductOrderUpdate {
	pou.mutation.ResetPrice()
	pou.mutation.SetPrice(i)
	return pou
}

// AddPrice adds i to the "price" field.
func (pou *ProductOrderUpdate) AddPrice(i int) *ProductOrderUpdate {
	pou.mutation.AddPrice(i)
	return pou
}

// SetCount sets the "count" field.
func (pou *ProductOrderUpdate) SetCount(i int) *ProductOrderUpdate {
	pou.mutation.ResetCount()
	pou.mutation.SetCount(i)
	return pou
}

// AddCount adds i to the "count" field.
func (pou *ProductOrderUpdate) AddCount(i int) *ProductOrderUpdate {
	pou.mutation.AddCount(i)
	return pou
}

// SetStartDate sets the "start_date" field.
func (pou *ProductOrderUpdate) SetStartDate(t time.Time) *ProductOrderUpdate {
	pou.mutation.SetStartDate(t)
	return pou
}

// SetEndDate sets the "end_date" field.
func (pou *ProductOrderUpdate) SetEndDate(t time.Time) *ProductOrderUpdate {
	pou.mutation.SetEndDate(t)
	return pou
}

// SetRegisteredAt sets the "registered_at" field.
func (pou *ProductOrderUpdate) SetRegisteredAt(t time.Time) *ProductOrderUpdate {
	pou.mutation.SetRegisteredAt(t)
	return pou
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (pou *ProductOrderUpdate) SetNillableRegisteredAt(t *time.Time) *ProductOrderUpdate {
	if t != nil {
		pou.SetRegisteredAt(*t)
	}
	return pou
}

// Mutation returns the ProductOrderMutation object of the builder.
func (pou *ProductOrderUpdate) Mutation() *ProductOrderMutation {
	return pou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pou *ProductOrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pou.hooks) == 0 {
		affected, err = pou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pou.mutation = mutation
			affected, err = pou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pou.hooks) - 1; i >= 0; i-- {
			mut = pou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pou *ProductOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := pou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pou *ProductOrderUpdate) Exec(ctx context.Context) error {
	_, err := pou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pou *ProductOrderUpdate) ExecX(ctx context.Context) {
	if err := pou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pou *ProductOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product_order.Table,
			Columns: product_order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product_order.FieldID,
			},
		},
	}
	if ps := pou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pou.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product_order.FieldName,
		})
	}
	if value, ok := pou.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldPrice,
		})
	}
	if value, ok := pou.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldPrice,
		})
	}
	if value, ok := pou.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldCount,
		})
	}
	if value, ok := pou.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldCount,
		})
	}
	if value, ok := pou.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldStartDate,
		})
	}
	if value, ok := pou.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldEndDate,
		})
	}
	if value, ok := pou.mutation.RegisteredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldRegisteredAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product_order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductOrderUpdateOne is the builder for updating a single Product_order entity.
type ProductOrderUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductOrderMutation
}

// SetName sets the "name" field.
func (pouo *ProductOrderUpdateOne) SetName(s string) *ProductOrderUpdateOne {
	pouo.mutation.SetName(s)
	return pouo
}

// SetPrice sets the "price" field.
func (pouo *ProductOrderUpdateOne) SetPrice(i int) *ProductOrderUpdateOne {
	pouo.mutation.ResetPrice()
	pouo.mutation.SetPrice(i)
	return pouo
}

// AddPrice adds i to the "price" field.
func (pouo *ProductOrderUpdateOne) AddPrice(i int) *ProductOrderUpdateOne {
	pouo.mutation.AddPrice(i)
	return pouo
}

// SetCount sets the "count" field.
func (pouo *ProductOrderUpdateOne) SetCount(i int) *ProductOrderUpdateOne {
	pouo.mutation.ResetCount()
	pouo.mutation.SetCount(i)
	return pouo
}

// AddCount adds i to the "count" field.
func (pouo *ProductOrderUpdateOne) AddCount(i int) *ProductOrderUpdateOne {
	pouo.mutation.AddCount(i)
	return pouo
}

// SetStartDate sets the "start_date" field.
func (pouo *ProductOrderUpdateOne) SetStartDate(t time.Time) *ProductOrderUpdateOne {
	pouo.mutation.SetStartDate(t)
	return pouo
}

// SetEndDate sets the "end_date" field.
func (pouo *ProductOrderUpdateOne) SetEndDate(t time.Time) *ProductOrderUpdateOne {
	pouo.mutation.SetEndDate(t)
	return pouo
}

// SetRegisteredAt sets the "registered_at" field.
func (pouo *ProductOrderUpdateOne) SetRegisteredAt(t time.Time) *ProductOrderUpdateOne {
	pouo.mutation.SetRegisteredAt(t)
	return pouo
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (pouo *ProductOrderUpdateOne) SetNillableRegisteredAt(t *time.Time) *ProductOrderUpdateOne {
	if t != nil {
		pouo.SetRegisteredAt(*t)
	}
	return pouo
}
// Mutation returns the ProductOrderMutation object of the builder.
func (pouo *ProductOrderUpdateOne) Mutation() *ProductOrderMutation {
	return pouo.mutation
}

// Save executes the query and returns the updated Product_order entity.
func (pouo *ProductOrderUpdateOne) Save(ctx context.Context) (*Product_order, error) {
	var (
		err  error
		node *Product_order
	)
	if len(pouo.hooks) == 0 {
		node, err = pouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pouo.mutation = mutation
			node, err = pouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pouo.hooks) - 1; i >= 0; i-- {
			mut = pouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pouo *ProductOrderUpdateOne) SaveX(ctx context.Context) *Product_order {
	node, err := pouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pouo *ProductOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := pouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pouo *ProductOrderUpdateOne) ExecX(ctx context.Context) {
	if err := pouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pouo *ProductOrderUpdateOne) sqlSave(ctx context.Context) (_node *Product_order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product_order.Table,
			Columns: product_order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product_order.FieldID,
			},
		},
	}
	id, ok := pouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product_order.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := pouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pouo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product_order.FieldName,
		})
	}
	if value, ok := pouo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldPrice,
		})
	}
	if value, ok := pouo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldPrice,
		})
	}
	if value, ok := pouo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldCount,
		})
	}
	if value, ok := pouo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldCount,
		})
	}
	if value, ok := pouo.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldStartDate,
		})
	}
	if value, ok := pouo.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldEndDate,
		})
	}
	if value, ok := pouo.mutation.RegisteredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldRegisteredAt,
		})
	}
	_node = &Product_order{config: pouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product_order.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
