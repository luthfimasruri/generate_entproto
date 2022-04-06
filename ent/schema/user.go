package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").Unique().Annotations(entproto.Field(2)),
		field.String("Phone").Unique().Annotations(entproto.Field(3)),
		field.String("Password").Sensitive().Annotations(entproto.Field(4)),
		field.String("FirstName").Annotations(entproto.Field(5)),
		field.String("LastName").Annotations(entproto.Field(6)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
