package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Departments holds the schema definition for the Departments entity.
type Departments struct {
	ent.Schema
}

// Fields of the Departments.
func (Departments) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("code"),
		field.String("name"),
		field.Bool("degree_type"),
		field.Int("quota"),
	}
}

// Edges of the Departments.
func (Departments) Edges() []ent.Edge {
	return nil
}
