package main

import (
	"net/http"

	"github.com/ecshreve/jepp/internal/ent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.ListenAndServe("localhost:3002", ent.ServeEntviz())
	// log.SetLevel(log.TraceLevel)
	// ss := jeppserver.NewServer()
	// ss.Start()
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
