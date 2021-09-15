package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("title").NotEmpty().MaxLen(75),
		field.Text("content").Unique().NotEmpty(),
		field.Int("views").Positive(),
		field.Int("likes").Positive(),
	}
}

// Edges of the Experience.
func (Experience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comments.Type),
		edge.From("user", User.Type).
			Ref("posts").
			Unique(),
	}
}
