package server

import (
	"strconv"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// registerUIHandlers registers route handlers for the UI.
func (s *Server) registerUIHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/server/templates/*")

	// s.Router.GET("/", s.BaseUIHandler)
	// s.Router.POST("/", s.BaseUIHandler)
	s.Router.GET("/:gameID/:categoryID/:clueID", s.ClueUIHandler)

	s.Router.GET("/debug", s.DebugUIHandler)

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) ClueUIHandler(c *gin.Context) {
	gameID, _ := strconv.ParseInt(c.Param("gameID"), 10, 64)
	categoryID, _ := strconv.ParseInt(c.Param("categoryID"), 10, 64)
	clueID, _ := strconv.ParseInt(c.Param("clueID"), 10, 64)

	clue, _ := s.DB.GetClue(clueID)
	game, _ := s.DB.GetGame(gameID)
	category, _ := s.DB.GetCategory(categoryID)
	clueJSON := s.jsonHelper(clue)

	debug := struct {
		*models.Clue
		*models.Game
		*models.Category
		*models.Stats
		ClueJSON string
	}{
		Clue:     clue,
		Game:     game,
		Category: category,
		Stats:    s.Stats,
		ClueJSON: pretty.Sprint(clueJSON),
	}
	c.HTML(200, "base.html.tpl", debug)
}

func (s *Server) PostFormHandler(trigger string, clueIdStr string) ([]*models.Clue, error) {
	clueId, _ := strconv.ParseInt(clueIdStr, 10, 64)
	clue, _ := s.DB.GetClue(clueId)
	var params models.CluesParams

	switch trigger {
	case "cat-roll":
		params = models.CluesParams{
			GameID: clue.GameID,
		}
	case "clue-roll":
		params = models.CluesParams{
			GameID:     clue.GameID,
			CategoryID: clue.CategoryID,
		}
	case "game-roll":
	default:
		params = models.CluesParams{}
	}

	clues, _ := s.DB.ListClues(params)
	return clues, nil
}

func (s *Server) jsonHelper(clue *models.Clue) map[string]interface{} {
	cat, _ := s.DB.GetCategory(clue.CategoryID)
	game, _ := s.DB.GetGame(clue.GameID)

	return map[string]interface{}{
		"clueID":       clue.ClueID,
		"categoryID":   clue.CategoryID,
		"categoryName": cat.Name,
		"gameID":       clue.GameID,
		"gameDate":     game.GameDate,
		"question":     clue.Question,
		"answer":       clue.Answer,
	}
}

func (s *Server) DebugUIHandler(c *gin.Context) {
	clue, _ := s.DB.GetRandomClue(nil)
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
