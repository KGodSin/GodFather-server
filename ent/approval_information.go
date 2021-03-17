// Code generated by entc, DO NOT EDIT.

package ent

import (
	"GodFather-server/ent/approval_information"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Approval_information is the model entity for the Approval_information schema.
type Approval_information struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status approval_information.Status `json:"status,omitempty"`
	// ApprovalDate holds the value of the "approval_date" field.
	ApprovalDate time.Time `json:"approval_date,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID string `json:"product_id,omitempty"`
	// RegisteredAt holds the value of the "registered_at" field.
	RegisteredAt                  string `json:"registered_at,omitempty"`
	product_approval_informations *int
	user_approval_informations    *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Approval_information) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case approval_information.FieldID, approval_information.FieldCount:
			values[i] = &sql.NullInt64{}
		case approval_information.FieldStatus, approval_information.FieldUserID, approval_information.FieldProductID, approval_information.FieldRegisteredAt:
			values[i] = &sql.NullString{}
		case approval_information.FieldApprovalDate:
			values[i] = &sql.NullTime{}
		case approval_information.ForeignKeys[0]: // product_approval_informations
			values[i] = &sql.NullInt64{}
		case approval_information.ForeignKeys[1]: // user_approval_informations
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Approval_information", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Approval_information fields.
func (ai *Approval_information) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case approval_information.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ai.ID = int(value.Int64)
		case approval_information.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ai.Status = approval_information.Status(value.String)
			}
		case approval_information.FieldApprovalDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field approval_date", values[i])
			} else if value.Valid {
				ai.ApprovalDate = value.Time
			}
		case approval_information.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				ai.Count = int(value.Int64)
			}
		case approval_information.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ai.UserID = value.String
			}
		case approval_information.FieldProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				ai.ProductID = value.String
			}
		case approval_information.FieldRegisteredAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registered_at", values[i])
			} else if value.Valid {
				ai.RegisteredAt = value.String
			}
		case approval_information.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_approval_informations", value)
			} else if value.Valid {
				ai.product_approval_informations = new(int)
				*ai.product_approval_informations = int(value.Int64)
			}
		case approval_information.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_approval_informations", value)
			} else if value.Valid {
				ai.user_approval_informations = new(int)
				*ai.user_approval_informations = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Approval_information.
// Note that you need to call Approval_information.Unwrap() before calling this method if this Approval_information
// was returned from a transaction, and the transaction was committed or rolled back.
func (ai *Approval_information) Update() *ApprovalInformationUpdateOne {
	return (&Approval_informationClient{config: ai.config}).UpdateOne(ai)
}

// Unwrap unwraps the Approval_information entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ai *Approval_information) Unwrap() *Approval_information {
	tx, ok := ai.config.driver.(*txDriver)
	if !ok {
		panic("ent: Approval_information is not a transactional entity")
	}
	ai.config.driver = tx.drv
	return ai
}

// String implements the fmt.Stringer.
func (ai *Approval_information) String() string {
	var builder strings.Builder
	builder.WriteString("Approval_information(")
	builder.WriteString(fmt.Sprintf("id=%v", ai.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ai.Status))
	builder.WriteString(", approval_date=")
	builder.WriteString(ai.ApprovalDate.Format(time.ANSIC))
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", ai.Count))
	builder.WriteString(", user_id=")
	builder.WriteString(ai.UserID)
	builder.WriteString(", product_id=")
	builder.WriteString(ai.ProductID)
	builder.WriteString(", registered_at=")
	builder.WriteString(ai.RegisteredAt)
	builder.WriteByte(')')
	return builder.String()
}

// Approval_informations is a parsable slice of Approval_information.
type Approval_informations []*Approval_information

func (ai Approval_informations) config(cfg config) {
	for _i := range ai {
		ai[_i].config = cfg
	}
}
