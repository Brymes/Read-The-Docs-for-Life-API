package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Comments holds the schema definition for the Comments entity.
type Comments struct {
	ent.Schema
}

// Fields of the Comments.
func (Comments) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Text("comment").Unique().NotEmpty(),
		field.Int("likes").Positive(),
	}
}

// Edges of the Comments.
func (Comments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Experience.Type).
			Ref("comments").
			Unique(),
		edge.From("user", User.Type).
			Ref("comments").
			Unique(),
	}
}
