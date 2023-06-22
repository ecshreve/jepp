package main

import (
	"github.com/ecshreve/jepp/docs"
	"github.com/ecshreve/jepp/pkg/server"
	log "github.com/sirupsen/logrus"
)

//	@contact.name	shreve
//	@contact.url	http://github.com/ecshreve
//	@contact.email	eric@shreve.dev

// @license.name	MIT
// @license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
func main() {
	docs.SwaggerInfo.Title = "Jepp API"
	docs.SwaggerInfo.Description = "This is a simple jeopardy server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "jepp.app"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"https"}

	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp API server...")

	srv := server.NewServer()
	srv.Serve()
}
