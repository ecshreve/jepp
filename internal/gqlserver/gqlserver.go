package gqlserver

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ecshreve/jepp/internal/ent"
	gqlserver "github.com/ecshreve/jepp/internal/gqlserver/resolvers"
	"github.com/ecshreve/jepp/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	client *ent.Client
	db     *sql.DB
	router *gin.Engine
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	cl, _ := utils.InitDB()
	h := handler.NewDefaultServer(gqlserver.NewSchema(cl))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func NewServer(router *gin.Engine) *Server {
	log.SetLevel(log.DebugLevel)
	cl, db := utils.InitDB()

	r := router
	if router == nil {
		r = gin.Default()
	}

	r.POST("/query", graphqlHandler())
	r.GET("/graphql", playgroundHandler())

	return &Server{
		client: cl,
		db:     db,
		router: r,
	}
}

func RunServer() {
	log.SetLevel(log.DebugLevel)
	r := gin.Default()

	NewServer(r)
	r.Use(cors.Default())
	log.Info("listening on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("http server terminated", err)
	}
}
