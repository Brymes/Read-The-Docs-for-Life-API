package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tags holds the schema definition for the Tags entity.
type Tags struct {
	ent.Schema
}

// Fields of the Tags.
func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").Unique().NotEmpty().MaxLen(25),
	}
}

// Edges of the Tags.
func (Tags) Edges() []ent.Edge {
	return nil
}
