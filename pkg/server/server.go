package server

import (
	"fmt"
	"log"
	"os"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/server/pagination"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server is the API server.
type Server struct {
	ID     string
	Router *gin.Engine
	Clock  clock.Clock
	DB     *models.JeppDB
	Stats  *models.Stats
}

// NewServer returns a new API server.
func NewServer() *Server {
	dbname := "jeppdb"
	dbuser := "jepp-user"
	dbpass := os.Getenv("DB_PASS")
	dbaddr := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	jdb := models.NewDB(dbname, dbuser, dbpass, dbaddr)
	stats, _ := jdb.GetStats()

	// Expose prometheus metrics on /metrics.
	r := gin.Default()
	// p := ginprometheus.NewPrometheus("gin")
	// p.SetListenAddressWithRouter(":9990", gin.New())
	// p.Use(r)

	s := &Server{
		ID:     "SERVER",
		Router: r,
		Clock:  clock.New(),
		DB:     jdb,
		Stats:  stats,
	}

	s.registerHandlers()

	if os.Getenv("JEPP_LOCAL_DEV") == "true" {
		s.registerDevHandlers()
	}

	return s
}

// registerHandlers registers all the handlers for the server.
func (s *Server) registerHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/server/templates/prod/*")

	s.Router.GET("/", func(ctx *gin.Context) { s.BaseUIHandler(ctx) })
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := s.Router.Group("/api")
	api.GET("/", s.BaseHandler)
	api.GET("/ping", s.PingHandler)

	listPaginator := pagination.New("page", "size", "1", "10", 1, 100)

	clue := api.Group("/clues")
	clue.GET("", listPaginator, s.CluesHandler)
	clue.GET("/random", s.RandomClueHandler)
	clue.GET("/:clueID", s.ClueHandler)

	games := api.Group("/games")
	games.GET("", listPaginator, s.GamesHandler)
	games.GET("/random", s.RandomGameHandler)
	games.GET("/:gameID", s.GameHandler)

	category := api.Group("/categories")
	category.GET("", listPaginator, s.CategoriesHandler)
	category.GET("/random", s.RandomCategoryHandler)
	category.GET("/:categoryID", s.CategoryHandler)

	// if err := s.Router.SetTrustedProxies(nil); err != nil {
	// 	log.Error(oops.Wrapf(err, "unable to set proxies"))
	// }
}

// Serve starts the server.
func (s *Server) Serve() error {
	log.Fatal(autotls.Run(s.Router, "jepp.app"))
	// err := s.Router.RunTLS(":8880", "pkg/server/certs/server.pem", "pkg/server/certs/server.key")
	// if err != nil {
	// 	return oops.Wrapf(err, "gin server returned error")
	// }
	return nil
}
