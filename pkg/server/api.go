package server

import (
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// registerHandlers registers all the handlers for the server.
func (s *Server) registerAPIHandlers() {
	api := s.Router.Group("/api")
	api.GET("/", s.BaseHandler)
	api.GET("/ping", s.PingHandler)

	s.registerGameHandlers(api)
	s.registerCategoryHandlers(api)
	s.registerClueHandlers(api)

	// This is not safe for production use but it's fine for playing
	// around locally.
	if err := s.Router.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}
}

// BaseHandler godoc
//
//	@Summary		Base handler
//	@Description	Show available endpoints
//
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ [get]
func (s *Server) BaseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api root",
		"available endpoints": []string{
			"/api",
			"/api/ping",
			"/api/games",
			"/api/games/random",
			"/api/games/:gameID",
			"/api/games/:gameID/api/clues",
			"/api/categories",
			"/api/categories/random",
			"/api/categories/:categoryID",
			"/api/categories/:categoryID/api/clues",
			"/api/clues",
			"/api/clues/random",
			"/api/clues/:clueID",
		},
	})
}

// PingHandler godoc
//
//	@Summary		Show the status of server
//	@Description	Get the status of server
//
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ping [get]
func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
