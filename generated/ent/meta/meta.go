// Code generated by ent, DO NOT EDIT.

package meta

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the meta type in the database.
	Label = "meta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeployment holds the string denoting the deployment field in the database.
	FieldDeployment = "deployment"
	// FieldHasIndexingErrors holds the string denoting the hasindexingerrors field in the database.
	FieldHasIndexingErrors = "has_indexing_errors"
	// EdgeBlock holds the string denoting the block edge name in mutations.
	EdgeBlock = "block"
	// Table holds the table name of the meta in the database.
	Table = "meta"
	// BlockTable is the table that holds the block relation/edge.
	BlockTable = "meta"
	// BlockInverseTable is the table name for the Block entity.
	// It exists in this package in order to avoid circular dependency with the "block" package.
	BlockInverseTable = "blocks"
	// BlockColumn is the table column denoting the block relation/edge.
	BlockColumn = "meta_block"
)

// Columns holds all SQL columns for meta fields.
var Columns = []string{
	FieldID,
	FieldDeployment,
	FieldHasIndexingErrors,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "meta"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"meta_block",
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

// OrderOption defines the ordering options for the Meta queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeployment orders the results by the deployment field.
func ByDeployment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeployment, opts...).ToFunc()
}

// ByHasIndexingErrors orders the results by the hasIndexingErrors field.
func ByHasIndexingErrors(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasIndexingErrors, opts...).ToFunc()
}

// ByBlockField orders the results by block field.
func ByBlockField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlockStep(), sql.OrderByField(field, opts...))
	}
}
func newBlockStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlockInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BlockTable, BlockColumn),
	)
}
