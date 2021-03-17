package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product_order holds the schema definition for the Product_order entity.
type Product_order struct {
	ent.Schema
}

// Fields of the Product_order.
func (Product_order) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("price"),
		field.Int("count"),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Time("registered_at").Default(time.Now),
	}
}

// Edges of the Product_order.
func (Product_order) Edges() []ent.Edge {
	return nil
}
