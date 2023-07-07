package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Annotations(
			entgql.OrderField("ID"),
		),
		field.Text("name").Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clues", Clue.Type).Annotations(
			entgql.RelayConnection(),
			entgql.Skip(),
			entoas.ListOperation(
				entoas.OperationPolicy(entoas.PolicyExpose),
			)),
	}
}

func (Category) Annotations() []schema.Annotation {
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
