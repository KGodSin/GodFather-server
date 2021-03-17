// Code generated by entc, DO NOT EDIT.

package ent

import (
	"GodFather-server/ent/product_order"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductOrderCreate is the builder for creating a Product_order entity.
type ProductOrderCreate struct {
	config
	mutation *ProductOrderMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (poc *ProductOrderCreate) SetName(s string) *ProductOrderCreate {
	poc.mutation.SetName(s)
	return poc
}

// SetPrice sets the "price" field.
func (poc *ProductOrderCreate) SetPrice(i int) *ProductOrderCreate {
	poc.mutation.SetPrice(i)
	return poc
}

// SetCount sets the "count" field.
func (poc *ProductOrderCreate) SetCount(i int) *ProductOrderCreate {
	poc.mutation.SetCount(i)
	return poc
}

// SetStartDate sets the "start_date" field.
func (poc *ProductOrderCreate) SetStartDate(t time.Time) *ProductOrderCreate {
	poc.mutation.SetStartDate(t)
	return poc
}

// SetEndDate sets the "end_date" field.
func (poc *ProductOrderCreate) SetEndDate(t time.Time) *ProductOrderCreate {
	poc.mutation.SetEndDate(t)
	return poc
}

// SetRegisteredAt sets the "registered_at" field.
func (poc *ProductOrderCreate) SetRegisteredAt(t time.Time) *ProductOrderCreate {
	poc.mutation.SetRegisteredAt(t)
	return poc
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (poc *ProductOrderCreate) SetNillableRegisteredAt(t *time.Time) *ProductOrderCreate {
	if t != nil {
		poc.SetRegisteredAt(*t)
	}
	return poc
}

// Mutation returns the ProductOrderMutation object of the builder.
func (poc *ProductOrderCreate) Mutation() *ProductOrderMutation {
	return poc.mutation
}

// Save creates the Product_order in the database.
func (poc *ProductOrderCreate) Save(ctx context.Context) (*Product_order, error) {
	var (
		err  error
		node *Product_order
	)
	poc.defaults()
	if len(poc.hooks) == 0 {
		if err = poc.check(); err != nil {
			return nil, err
		}
		node, err = poc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = poc.check(); err != nil {
				return nil, err
			}
			poc.mutation = mutation
			node, err = poc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(poc.hooks) - 1; i >= 0; i-- {
			mut = poc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, poc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (poc *ProductOrderCreate) SaveX(ctx context.Context) *Product_order {
	v, err := poc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (poc *ProductOrderCreate) defaults() {
	if _, ok := poc.mutation.RegisteredAt(); !ok {
		v := product_order.DefaultRegisteredAt()
		poc.mutation.SetRegisteredAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (poc *ProductOrderCreate) check() error {
	if _, ok := poc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := poc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := poc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New("ent: missing required field \"count\"")}
	}
	if _, ok := poc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New("ent: missing required field \"start_date\"")}
	}
	if _, ok := poc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New("ent: missing required field \"end_date\"")}
	}
	if _, ok := poc.mutation.RegisteredAt(); !ok {
		return &ValidationError{Name: "registered_at", err: errors.New("ent: missing required field \"registered_at\"")}
	}
	return nil
}

func (poc *ProductOrderCreate) sqlSave(ctx context.Context) (*Product_order, error) {
	_node, _spec := poc.createSpec()
	if err := sqlgraph.CreateNode(ctx, poc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (poc *ProductOrderCreate) createSpec() (*Product_order, *sqlgraph.CreateSpec) {
	var (
		_node = &Product_order{config: poc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product_order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product_order.FieldID,
			},
		}
	)
	if value, ok := poc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product_order.FieldName,
		})
		_node.Name = value
	}
	if value, ok := poc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := poc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product_order.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := poc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := poc.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := poc.mutation.RegisteredAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product_order.FieldRegisteredAt,
		})
		_node.RegisteredAt = value
	}
	return _node, _spec
}

// ProductOrderCreateBulk is the builder for creating many Product_order entities in bulk.
type ProductOrderCreateBulk struct {
	config
	builders []*ProductOrderCreate
}

// Save creates the Product_order entities in the database.
func (pocb *ProductOrderCreateBulk) Save(ctx context.Context) ([]*Product_order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pocb.builders))
	nodes := make([]*Product_order, len(pocb.builders))
	mutators := make([]Mutator, len(pocb.builders))
	for i := range pocb.builders {
		func(i int, root context.Context) {
			builder := pocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductOrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pocb *ProductOrderCreateBulk) SaveX(ctx context.Context) []*Product_order {
	v, err := pocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
