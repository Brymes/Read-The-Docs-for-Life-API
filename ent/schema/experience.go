package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Experience holds the schema definition for the Experience entity.
type Experience struct {
	ent.Schema
}

// Fields of the Experience.
func (Experience) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.String("Title").NotEmpty().MaxLen(75),
		field.Text("Content").Unique().NotEmpty(),
		field.Int("Views").Positive(),
	}
}

// Edges of the Experience.
func (Experience) Edges() []ent.Edge {
	return nil
}
