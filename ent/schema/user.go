package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.String("lastName").MaxLen(75).Optional(),
		field.String("firstName").MaxLen(75).Optional(),
		field.Int("age").Positive().Optional(),
		field.String("profile").Optional(),
		field.String("email").MaxLen(40).NotEmpty(), //. TODO Validate(),
		field.Text("password").NotEmpty().Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comments.Type),
		edge.To("posts", Experience.Type),
	}
}
