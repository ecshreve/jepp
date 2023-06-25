package main

import (
	"github.com/ecshreve/jepp/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp API server...")

	srv := server.NewServer()
	if err := srv.Router.Run(":8880"); err != nil {
		log.Fatal(err)
	}
}
