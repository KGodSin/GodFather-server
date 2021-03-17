package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Approval_information holds the schema definition for the Approval_information entity.
type Approval_information struct {
	ent.Schema
}

// Fields of the Approval_information.
func (Approval_information) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("approved", "not_approved").Default("not_approved"),
		field.Time("approval_date"),
		field.Int("count"),
		field.String("user_id"),
		field.String("product_id"),
		field.Time("registered_at").Default(time.Now),
	}
}

// Edges of the Approval_information.
func (Approval_information) Edges() []ent.Edge {
	return nil
}
