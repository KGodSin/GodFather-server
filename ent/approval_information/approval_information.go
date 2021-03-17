// Code generated by entc, DO NOT EDIT.

package approval_information

import (
	"fmt"
)

const (
	// Label holds the string label denoting the approval_information type in the database.
	Label = "approval_information"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldApprovalDate holds the string denoting the approval_date field in the database.
	FieldApprovalDate = "approval_date"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// Table holds the table name of the approval_information in the database.
	Table = "approval_informations"
)

// Columns holds all SQL columns for approval_information fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldApprovalDate,
	FieldCount,
	FieldUserID,
	FieldProductID,
	FieldRegisteredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "approval_informations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_approval_informations",
	"user_approval_informations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// StatusNotApproved is the default value of the Status enum.
const DefaultStatus = StatusNotApproved

// Status values.
const (
	StatusApproved    Status = "approved"
	StatusNotApproved Status = "not_approved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusApproved, StatusNotApproved:
		return nil
	default:
		return fmt.Errorf("approval_information: invalid enum value for status field: %q", s)
	}
}
