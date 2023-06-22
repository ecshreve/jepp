package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// registerUIHandlers registers route handlers for the UI.
func (s *Server) registerUIHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/server/templates/prod/*")

	s.Router.GET("/", s.BaseUIHandler)
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// BaseUIHandler handles the base UI route.
func (s *Server) BaseUIHandler(c *gin.Context) {
	clue, err := s.DB.GetRandomClue(nil)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch random clue"})
		return
	}

	cat, err := s.DB.GetCategory(clue.CategoryID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch category for clue"})
		return
	}

	game, err := s.DB.GetGame(clue.GameID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch game for clue"})
		return
	}

	clueJSON, _ := json.Marshal(clue)

	c.HTML(200, "base.html.tpl", gin.H{
		"Stats":    s.Stats,
		"Clue":     clue,
		"ClueJSON": string(clueJSON),
		"Game":     game,
		"Category": cat,
	})
}
