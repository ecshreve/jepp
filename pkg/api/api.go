package api

import (
	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	ID     string
	Router *gin.Engine
	Clock  clock.Clock
	DB     *models.JeppDB
}

func NewServer() *Server {
	s := &Server{
		ID:     "SERVER",
		Router: gin.Default(),
		Clock:  clock.New(),
		DB:     models.NewDB(),
	}

	s.registerHandlers()

	return s
}

// CluesForGameHandler returns a list of clues for a given game.
func (s *Server) CluesForGameHandler(c *gin.Context) {
	gameID := c.Param("gameID")
	clues, err := s.DB.GetCluesForGame(gameID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues for game %s", gameID))
		c.JSON(500, gin.H{
			"message": "unable to get clues",
		})
		return
	}

	c.JSON(200, gin.H{
		"game_id": gameID,
		"clues":   clues,
	})
}

// GamesHandler returns a list of games.
func (s *Server) GamesHandler(c *gin.Context) {
	games, err := s.DB.GetAllGames()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		c.JSON(500, gin.H{
			"message": "unable to get games",
		})
		return
	}

	c.JSON(200, gin.H{
		"games": games,
	})
}

func (s *Server) BaseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api root",
		"available endpoints": []string{
			"/",
			"/ping",
			"/games",
		},
	})
}

func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) registerHandlers() {
	s.Router.GET("/", s.BaseHandler)
	s.Router.GET("/ping", s.PingHandler)
	s.Router.GET("/games", s.GamesHandler)
	s.Router.GET("/games/:gameID/clues", s.CluesForGameHandler)

	if err := s.Router.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}
}

func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
