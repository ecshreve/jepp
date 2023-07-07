package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Annotations(
			entgql.OrderField("ID"),
		),
		field.Int("number").Annotations(
			entgql.OrderField("NUMBER"),
		),
		field.Time("startDate").Annotations(
			entgql.OrderField("START_DATE"),
		),
		field.Time("endDate").Annotations(
			entgql.OrderField("END_DATE"),
		),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("games", Game.Type).Annotations(
			entgql.RelayConnection(),
			entgql.Skip(),
			entoas.ListOperation(
				entoas.OperationPolicy(entoas.PolicyExpose),
			),
		),
	}
}

func (Season) Annotations() []schema.Annotation {
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
