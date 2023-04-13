// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"pamukkale_university/ent/departments"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Departments is the model entity for the Departments schema.
type Departments struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DegreeType holds the value of the "degree_type" field.
	DegreeType bool `json:"degree_type,omitempty"`
	// Quota holds the value of the "quota" field.
	Quota        int `json:"quota,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Departments) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case departments.FieldDegreeType:
			values[i] = new(sql.NullBool)
		case departments.FieldQuota:
			values[i] = new(sql.NullInt64)
		case departments.FieldCode, departments.FieldName:
			values[i] = new(sql.NullString)
		case departments.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Departments fields.
func (d *Departments) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case departments.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case departments.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				d.Code = value.String
			}
		case departments.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case departments.FieldDegreeType:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field degree_type", values[i])
			} else if value.Valid {
				d.DegreeType = value.Bool
			}
		case departments.FieldQuota:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quota", values[i])
			} else if value.Valid {
				d.Quota = int(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Departments.
// This includes values selected through modifiers, order, etc.
func (d *Departments) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this Departments.
// Note that you need to call Departments.Unwrap() before calling this method if this Departments
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Departments) Update() *DepartmentsUpdateOne {
	return NewDepartmentsClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Departments entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Departments) Unwrap() *Departments {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Departments is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Departments) String() string {
	var builder strings.Builder
	builder.WriteString("Departments(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("code=")
	builder.WriteString(d.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("degree_type=")
	builder.WriteString(fmt.Sprintf("%v", d.DegreeType))
	builder.WriteString(", ")
	builder.WriteString("quota=")
	builder.WriteString(fmt.Sprintf("%v", d.Quota))
	builder.WriteByte(')')
	return builder.String()
}

// DepartmentsSlice is a parsable slice of Departments.
type DepartmentsSlice []*Departments
