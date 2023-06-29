package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ecshreve/jepp/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// package main

// import (
// 	"os"

// 	"github.com/ecshreve/jepp/docs"
// 	"github.com/ecshreve/jepp/pkg/server"
// 	log "github.com/sirupsen/logrus"
// )

// // @title			Jepp API Documentation
// // @description	This is a simple api to access jeopardy data.
// // @version		1.0
// // @BasePath		/api
// //
// // @contact.name	shreve
// //
// // @license.name	MIT License
// // @license.url	https://github.com/ecshreve/jepp/blob/main/LICENSE
// func main() {
// 	if os.Getenv("JEPP_ENV") == "prod" {
// 		docs.SwaggerInfo.Host = "jepp.app"
// 	}

// 	log.SetLevel(log.DebugLevel)
// 	log.Info("Starting Jepp API server...")

// 	srv := server.NewServer()
// 	if err := srv.Router.Run(":8880"); err != nil {
// 		log.Fatal(err)
// 	}
// }
