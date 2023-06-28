package server

import (
	"github.com/benbjohnson/clock"
	_ "github.com/ecshreve/jepp/docs"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *models.JeppDB

// Server is the API server.
type Server struct {
	ID     string
	Clock  clock.Clock
	Router *gin.Engine
}

// NewServer returns a new API server.
func NewServer(jdb *models.JeppDB) *Server {
	s := &Server{
		ID:    "SERVER",
		Clock: clock.New(),
	}
	db = jdb
	s.Router = registerHandlers()

	// TODO: fix this
	// if os.Getenv("JEPP_ENV") == "dev" {
	// 	registerDevHandlers()
	// }

	log.Infof("Server %#v created", s)
	return s
}

type Filter struct {
	Random *bool  `form:"random"`
	ID     *int64 `form:"id"`
	Page   *int64 `form:"page,default=0" binding:"min=0"`
	Limit  *int64 `form:"limit,default=1" binding:"min=1,max=100"`
}

func registerHandlers() *gin.Engine {
	// Explicitly setting to debug mode, surfaces extra logging.
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.StaticFile("style.css", "./static/site/style.css")
	r.StaticFile("favicon.ico", "./static/site/favicon.ico")
	r.LoadHTMLGlob("pkg/server/templates/prod/*")

	r.GET("/", BaseUIHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	api.GET("/clue", ClueHandler)
	api.GET("/game", GameHandler)
	api.GET("/category", CategoryHandler)

	// Basic health check endpoint.
	// TODO: pull into isolated handler with docs
	api.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}

	return r
}
