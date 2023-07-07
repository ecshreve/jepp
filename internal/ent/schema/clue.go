package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Clue holds the schema definition for the Clue entity.
type Clue struct {
	ent.Schema
}

// Fields of the Clue.
func (Clue) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Annotations(
			entproto.Field(1),
			entgql.OrderField("ID"),
		),
		field.Text("question").Annotations(
			entproto.Field(2),
			entgql.OrderField("QUESTION"),
		),
		field.Text("answer").Annotations(
			entproto.Field(3),
			entgql.OrderField("ANSWER"),
		),
		field.Int("category_id").Annotations(
			entproto.Field(4),
			entgql.Skip(),
		),
		field.Int("game_id").Annotations(
			entproto.Field(5),
			entgql.Skip(),
		),
	}
}

// Edges of the Game.
func (Clue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("clues").
			Field("category_id").
			Unique().
			Required().Annotations(
			entproto.Field(6),
			entgql.OrderField("CATEGORY_NAME"),
		),
		edge.From("game", Game.Type).
			Ref("clues").
			Field("game_id").
			Unique().
			Required().Annotations(
			entproto.Field(7),
		),
	}
}

func (Clue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
		entgql.RelayConnection(),
		entoas.ReadOperation(
			entoas.OperationPolicy(entoas.PolicyExpose),
		),
		entoas.ListOperation(
			entoas.OperationPolicy(entoas.PolicyExpose),
		),
	}
}
