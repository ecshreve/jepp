//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/Yamashou/gqlgenc/clientgen"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	queries := []string{"client.query", "fragemt.query"}
	clientPackage := config.PackageConfig{
		Filename: "graph/generated/client.go",
		Package:  "graph",
	}

	clientPlugin := clientgen.New(queries, clientPackage, nil)
	err = api.Generate(cfg,
		api.AddPlugin(clientPlugin),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
