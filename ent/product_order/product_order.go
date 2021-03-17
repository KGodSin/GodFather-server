// Code generated by entc, DO NOT EDIT.

package product_order

import (
	"time"
)
const (
	// Label holds the string label denoting the product_order type in the database.
	Label = "product_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// Table holds the table name of the product_order in the database.
	Table = "product_orders"
)

// Columns holds all SQL columns for product_order fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPrice,
	FieldCount,
	FieldStartDate,
	FieldEndDate,
	FieldRegisteredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "product_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_product_orders",
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

var (
	// DefaultRegisteredAt holds the default value on creation for the "registered_at" field.
	DefaultRegisteredAt func() time.Time
)
