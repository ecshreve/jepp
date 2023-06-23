package server

import (
	"os"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// Server is the API server.
type Server struct {
	ID     string
	Clock  clock.Clock
	Router *gin.Engine
	JDB    *sqlx.DB
}

// NewServer returns a new API server.
func NewServer() *Server {
	s := &Server{
		ID:    "SERVER",
		Clock: clock.New(),
		JDB:   models.GetDBHandle(),
	}
	s.Router = registerHandlers()

	// TODO: fix this
	if os.Getenv("JEPP_LOCAL_DEV") == "true" {
		registerDevHandlers()
	}

	log.Infof("Server %#v created", s)
	return s
}

func registerHandlers() *gin.Engine {
	r := gin.Default()
	r.StaticFile("style.css", "./static/style.css")
	r.StaticFile("favicon.ico", "./static/favicon.ico")
	r.StaticFile("swagger.json", "./docs/swagger.json")

	r.LoadHTMLGlob("pkg/server/templates/prod/*")

	r.GET("/ui", BaseUIHandler)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	api.GET("", BaseHandler)
	api.GET("/ping", PingHandler)

	api.GET("/clues", CluesHandler)
	api.GET("/games", GamesHandler)
	api.GET("/categories", CategoriesHandler)

	api.GET("/random/game", RandomGameHandler)
	api.GET("/random/category", RandomCategoryHandler)
	api.GET("/random/clue", RandomClueHandler)

	// if err := r.SetTrustedProxies(nil); err != nil {
	// 	log.Error(oops.Wrapf(err, "unable to set proxies"))
	// }

	return r
}
