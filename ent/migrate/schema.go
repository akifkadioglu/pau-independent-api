// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "degree_type", Type: field.TypeBool},
		{Name: "quota", Type: field.TypeInt},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
	}
)

func init() {
}
