package server

import (
	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
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
	jdb := models.NewDB()
	stats, _ := jdb.GetStats()

	s := &Server{
		ID:     "SERVER",
		Router: gin.Default(),
		Clock:  clock.New(),
		DB:     jdb,
		Stats:  stats,
		QZ: &QuizSession{
			Clues:     make([]*models.Clue, 0),
			Correct:   0,
			Incorrect: 0,
			Total:     0,
		},
	}

	s.registerAPIHandlers()
	s.registerUIHandlers()

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
