package main

import (
	"os"

	"github.com/ecshreve/jepp/docs"
	"github.com/ecshreve/jepp/pkg/server"
	log "github.com/sirupsen/logrus"
)

//	@contact.name	shreve
//	@contact.url	http://github.com/ecshreve
//	@contact.email	eric@shreve.dev

//	@license.name	MIT
//	@license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
func main() {
	docs.SwaggerInfo.Title = "Jepp API Documentation"
	docs.SwaggerInfo.Description = "This is a simple api to access jeopardy data."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("API_HOST")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"https"}

	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp API server...")

	srv := server.NewServer()
	srv.Serve()
}
