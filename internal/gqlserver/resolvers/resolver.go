package gqlserver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/ecshreve/jepp/internal/ent"
	gqlserver "github.com/ecshreve/jepp/internal/gqlserver/gen"
	_ "github.com/hedwigz/entviz"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gqlserver.NewExecutableSchema(gqlserver.Config{
		Resolvers: &Resolver{client},
	})
}
