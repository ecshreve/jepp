package main

import "github.com/ecshreve/jepp/pkg/api"

func main() {
	apiServer := api.NewServer()
	apiServer.Serve()
}
