package server

import (
	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
)

// Server is the API server.
type Server struct {
	ID     string
	Router *gin.Engine
	Clock  clock.Clock
	DB     *models.JeppDB
}

// NewServer returns a new API server.
func NewServer() *Server {
	s := &Server{
		ID:     "SERVER",
		Router: gin.Default(),
		Clock:  clock.New(),
		DB:     models.NewDB(),
	}

	s.registerAPIHandlers()
	s.registerUIHandlers()

	return s
}

// Serve starts the server.
func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
