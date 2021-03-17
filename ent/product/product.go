// Code generated by entc, DO NOT EDIT.

package product

import (
	"time"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// EdgeProductOrders holds the string denoting the product_orders edge name in mutations.
	EdgeProductOrders = "product_orders"
	// EdgeApprovalInformations holds the string denoting the approval_informations edge name in mutations.
	EdgeApprovalInformations = "approval_informations"
	// Table holds the table name of the product in the database.
	Table = "products"
	// ProductOrdersTable is the table the holds the product_orders relation/edge.
	ProductOrdersTable = "product_orders"
	// ProductOrdersInverseTable is the table name for the Product_order entity.
	// It exists in this package in order to avoid circular dependency with the "product_order" package.
	ProductOrdersInverseTable = "product_orders"
	// ProductOrdersColumn is the table column denoting the product_orders relation/edge.
	ProductOrdersColumn = "product_product_orders"
	// ApprovalInformationsTable is the table the holds the approval_informations relation/edge.
	ApprovalInformationsTable = "approval_informations"
	// ApprovalInformationsInverseTable is the table name for the Approval_information entity.
	// It exists in this package in order to avoid circular dependency with the "approval_information" package.
	ApprovalInformationsInverseTable = "approval_informations"
	// ApprovalInformationsColumn is the table column denoting the approval_informations relation/edge.
	ApprovalInformationsColumn = "product_approval_informations"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRegisteredAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRegisteredAt holds the default value on creation for the "registered_at" field.
	DefaultRegisteredAt func() time.Time
)
