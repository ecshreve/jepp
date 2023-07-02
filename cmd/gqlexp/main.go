package main

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
	"github.com/ecshreve/jepp/pkg/gqlclient/gen"
	"github.com/kr/pretty"
)

func main() {
	ctx := context.Background()

	client := &gen.Client{
		Client: client.NewClient(http.DefaultClient, "http://localhost:8880/gql/query"),
	}

	x, _ := client.GetSeasons(ctx)
	pretty.Print(x)
}
