package server

import (
	"fmt"
	"os"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type QuizSession struct {
	Clues     []*models.Clue
	Correct   int64
	Incorrect int64
	Total     int64
}

// Server is the API server.
type Server struct {
	ID     string
	Router *gin.Engine
	Clock  clock.Clock
	DB     *models.JeppDB
	Stats  *models.Stats
	QZ     *QuizSession
}

// NewServer returns a new API server.
func NewServer() *Server {
	dbname := "jeppdb"
	dbuser := "jepp-user"
	dbpass := os.Getenv("DB_PASS")
	dbaddr := fmt.Sprintf("%s:3306", os.Getenv("DB_HOST"))

	jdb := models.NewDB(dbname, dbuser, dbpass, dbaddr)
	stats, _ := jdb.GetStats()

	// Expose prometheus metrics on /metrics.
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	s := &Server{
		ID:     "SERVER",
		Router: r,
		Clock:  clock.New(),
		DB:     jdb,
		Stats:  stats,
	}

	s.registerAPIHandlers()
	s.registerUIHandlers()

	if os.Getenv("JEPP_LOCAL_DEV") == "true" {
		s.registerDevHandlers()
	}

	return s
}

// Serve starts the server.
func (s *Server) Serve() error {
	// err := s.Router.Run(":8880")
	err := s.Router.RunTLS(":8880", "pkg/server/certs/server.pem", "pkg/server/certs/server.key")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
