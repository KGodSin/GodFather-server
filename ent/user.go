// Code generated by entc, DO NOT EDIT.

package ent

import (
	"GodFather-server/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// RegisteredAt holds the value of the "registered_at" field.
	RegisteredAt time.Time `json:"registered_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// ApprovalInformations holds the value of the approval_informations edge.
	ApprovalInformations []*Approval_information `json:"approval_informations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApprovalInformationsOrErr returns the ApprovalInformations value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ApprovalInformationsOrErr() ([]*Approval_information, error) {
	if e.loadedTypes[0] {
		return e.ApprovalInformations, nil
	}
	return nil, &NotLoadedError{edge: "approval_informations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = &sql.NullInt64{}
		case user.FieldUserID, user.FieldName, user.FieldPassword, user.FieldRole:
			values[i] = &sql.NullString{}
		case user.FieldRegisteredAt, user.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldRegisteredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field registered_at", values[i])
			} else if value.Valid {
				u.RegisteredAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryApprovalInformations queries the "approval_informations" edge of the User entity.
func (u *User) QueryApprovalInformations() *ApprovalInformationQuery {
	return (&UserClient{config: u.config}).QueryApprovalInformations(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(u.UserID)
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", registered_at=")
	builder.WriteString(u.RegisteredAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}