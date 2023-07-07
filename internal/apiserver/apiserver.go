package apiserver

import (
	"database/sql"

	"github.com/ecshreve/jepp/internal/ent"
	"github.com/ecshreve/jepp/internal/ent/ogent"
	"github.com/ecshreve/jepp/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

// serve® wraps the generated ogent.OgentHandler and overrides / adds http.Handler methods.
type Server struct {
	*ogent.OgentHandler
	db     *sql.DB
	client *ent.Client
	router *gin.Engine
}

var drv *sql.DB
var cl *ent.Client

func apiHandler() gin.HandlerFunc {
	// Start listening.
	srv, err := ogent.NewServer(Server{
		OgentHandler: ogent.NewOgentHandler(cl),
		db:           drv,
		client:       cl,
	})
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterAPIRoutes(api *gin.Engine) {
	api.GET("/games", apiHandler())
	api.GET("/games/:id", apiHandler())
	api.GET("/games/:id/clues", apiHandler())
	api.GET("/games/:id/season", apiHandler())

	api.GET("/categories", apiHandler())
	api.GET("/categories/:id", apiHandler())
	api.GET("/categories/:id/clues", apiHandler())

	api.GET("/clues", apiHandler())
	api.GET("/clues/:id", apiHandler())
	api.GET("/clues/:id/game", apiHandler())
	api.GET("/clues/:id/category", apiHandler())

	api.GET("/seasons", apiHandler())
	api.GET("/seasons/:id", apiHandler())
	api.GET("/seasons/:id/games", apiHandler())
}

func NewServer(router *gin.Engine) *Server {
	log.SetLevel(log.DebugLevel)
	cl, drv = utils.InitDB()

	r := router
	if router == nil {
		r = gin.Default()
	}

	RegisterAPIRoutes(r)
	return &Server{
		OgentHandler: ogent.NewOgentHandler(cl),
		db:           drv,
		client:       cl,
		router:       r,
	}
}

func RunServer() {
	log.SetLevel(log.DebugLevel)
	r := gin.Default()

	NewServer(r)
	r.Use(cors.Default())
	log.Info("listening on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("http server terminated", err)
	}
}
