package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
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
			entgql.OrderField("ID"),
		),
		field.Int("show").Annotations(
			entgql.OrderField("SHOW"),
		),
		field.Time("airDate").Annotations(
			entgql.OrderField("AIR_DATE"),
		),
		field.Time("tapeDate").Annotations(
			entgql.OrderField("TAPE_DATE"),
		),
		field.Int("season_id").Annotations(
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
			Required(),
		edge.To("clues", Clue.Type).Annotations(
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
		entgql.RelayConnection(),
		entoas.ReadOperation(
			entoas.OperationPolicy(entoas.PolicyExpose),
		),
		entoas.ListOperation(
			entoas.OperationPolicy(entoas.PolicyExpose),
		),
	}
}
