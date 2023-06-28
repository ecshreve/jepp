package main

import (
	"os"

	"github.com/ecshreve/jepp/docs"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/server"
	log "github.com/sirupsen/logrus"
)

// @title			Jepp API Documentation
// @description	This is a simple api to access jeopardy data.
// @version		1.0
// @BasePath		/api
//
// @contact.name	shreve
//
// @license.name	MIT License
// @license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
func main() {
	if os.Getenv("JEPP_ENV") == "prod" {
		docs.SwaggerInfo.Host = "jepp.app"
	}

	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp API server...")

	jdb := models.NewJeppDB("data/sqlite/jepp.db")

	srv := server.NewServer(jdb)
	if err := srv.Router.Run(":8880"); err != nil {
		log.Fatal(err)
	}
}
