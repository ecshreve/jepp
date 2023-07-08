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

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Annotations(
			entproto.Field(1),
			entgql.OrderField("ID"),
		),
		field.Int("show").Annotations(
			entproto.Field(2),
			entgql.OrderField("SHOW"),
		),
		field.Time("airDate").Annotations(
			entproto.Field(3),
			entgql.OrderField("AIR_DATE"),
		),
		field.Time("tapeDate").Annotations(
			entproto.Field(4),
			entgql.OrderField("TAPE_DATE"),
		),
		field.Int("season_id").Annotations(
			entproto.Field(5),
			entgql.Skip(),
		),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("games").
			Field("season_id").
			Unique().
			Required().Annotations(
			entproto.Field(6),
		),
		edge.To("clues", Clue.Type).Annotations(
			entproto.Field(7),

			entgql.RelayConnection(),
			entgql.OrderField("CLUE_ID"),
			entgql.Skip(),

			entoas.ListOperation(
				entoas.OperationPolicy(entoas.PolicyExpose),
			),
		),
	}
}

func (Game) Annotations() []schema.Annotation {
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
