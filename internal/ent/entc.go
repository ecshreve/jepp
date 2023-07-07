//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("gqlschema/ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(false),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec),
		entoas.DefaultPolicy(entoas.PolicyExclude),
		entoas.Mutations(func(_ *gen.Graph, spec *ogen.Spec) error {
			spec.Info.SetTitle("Jepp API").
				SetDescription("Jepp API").
				SetVersion("0.0.1")

			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogentext, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	ep, err := entproto.NewExtension()
	if err != nil {
		log.Fatalf("creating entproto extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ogentext, oas, ex, ep),
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
