// Code generated by ent, DO NOT EDIT.

package ent

import (
	"pamukkale_university/ent/departments"
	"pamukkale_university/ent/schema"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	departmentsFields := schema.Departments{}.Fields()
	_ = departmentsFields
	// departmentsDescID is the schema descriptor for id field.
	departmentsDescID := departmentsFields[0].Descriptor()
	// departments.DefaultID holds the default value on creation for the id field.
	departments.DefaultID = departmentsDescID.Default.(func() uuid.UUID)
}