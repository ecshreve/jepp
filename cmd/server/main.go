package main

import (
	"os"

	"github.com/ecshreve/jepp/docs"
	"github.com/ecshreve/jepp/pkg/server"
	log "github.com/sirupsen/logrus"
)

//	@title			Jepp API Documentation
//	@description	This is a simple api to access jeopardy data.
//	@version		1.0
//	@BasePath		/api
//
//	@contact.name	shreve
//
//	@license.name	MIT License
//	@license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
func main() {
	docs.SwaggerInfo.Host = os.Getenv("API_HOST")

	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp API server...")

	srv := server.NewServer()
	if err := srv.Router.Run(":8880"); err != nil {
		log.Fatal(err)
	}
}
