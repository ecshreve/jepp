package main

import (
	"github.com/ecshreve/jepp/internal/jeppserver"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)
	ss := jeppserver.NewServer()
	ss.Start()
	// log.SetLevel(log.DebugLevel)
	// r := gin.Default()

	// jeppui.NewUI(r)

	// // Register routes and initialize servers.
	// apiserver.NewServer(r)
	// gqlserver.NewServer(r)

	// // Run the ui.
	// log.Info("listening on :8082")
	// if err := r.Run(":8082"); err != nil {
	// 	log.Fatal("http server terminated", err)
	// }
}
