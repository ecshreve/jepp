package server

import (
	"context"
	"fmt"
	"os"

	"github.com/benbjohnson/clock"
	_ "github.com/ecshreve/jepp/docs"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server is the API server.
type Server struct {
	ID     string
	Clock  clock.Clock
	Router *gin.Engine
}

// NewServer returns a new API server.
func NewServer() *Server {
	s := &Server{
		ID:    "SERVER",
		Clock: clock.New(),
	}
	s.Router = registerHandlers()

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
	r.Use(GinContextToContextMiddleware())

	r.StaticFile("style.css", "./static/site/style.css")
	r.StaticFile("favicon.ico", "./static/site/favicon.ico")

	if os.Getenv("JEPP_ENV") != "test" {
		r.LoadHTMLGlob("pkg/server/templates/prod/*")

		r.GET("/", BaseUIHandler)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	gql := r.Group("/gql")
	gql.POST("/query", graphqlHandler())
	gql.GET("/", playgroundHandler())

	// api := r.Group("/api")
	// api.GET("/clue", ClueHandler)
	// api.GET("/game", GameHandler)
	// api.GET("/category", CategoryHandler)

	// Basic health check endpoint.
	// TODO: pull into isolated handler with docs
	// api.GET("/status", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"status": "ok"})
	// })

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}

	return r
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
