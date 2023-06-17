package server

import (
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// registerUIHandlers registers route handlers for the UI.
func (s *Server) registerUIHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")

	s.Router.LoadHTMLGlob("pkg/server/templates/*")

	s.Router.GET("/", s.RandomUIHandler)
	s.Router.GET("/debug", s.DebugUIHandler)

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) RandomUIHandler(c *gin.Context) {
	clue, _ := s.DB.GetRandomClue()
	game, _ := s.DB.GetGame(clue.GameID)
	category, _ := s.DB.GetCategory(clue.CategoryID)

	random := struct {
		*models.Clue
		*models.Game
		*models.Category
	}{
		Clue:     clue,
		Game:     game,
		Category: category,
	}
	c.HTML(200, "random.html.tpl", random)
}

func (s *Server) DebugUIHandler(c *gin.Context) {
	clue, _ := s.DB.GetRandomClue()
	game, _ := s.DB.GetGame(clue.GameID)
	category, _ := s.DB.GetCategory(clue.CategoryID)

	debug := struct {
		*models.Clue
		*models.Game
		*models.Category
	}{
		Clue:     clue,
		Game:     game,
		Category: category,
	}
	c.HTML(200, "debug.html.tpl", debug)
}
