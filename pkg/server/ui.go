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
	s.Router.GET("/:clueID", s.ClueUIHandler)
	s.Router.POST("/", s.ClueUIPOSTHandler)

	s.Router.GET("/debug", s.DebugUIHandler)

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

// ClueUIPOSTHandler godoc
func (s *Server) ClueUIPOSTHandler(c *gin.Context) {
	clueIDStr := c.PostForm("clue-sel")
	c.Params = gin.Params{gin.Param{Key: "clueID", Value: clueIDStr}}
	c.Request.Method = "GET"
	c.Redirect(302, "/"+clueIDStr)
}

// ClueUIHandler godoc
func (s *Server) ClueUIHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		s.ClueUIPOSTHandler(c)
		return
	}

	clueID, err := strconv.ParseInt(c.Param("clueID"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid gameID, categoryID, or clueID"})
		return
	}

	clue, err := s.DB.GetClue(clueID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid gameID, categoryID, or clueID"})
		return
	}

	game, err := s.DB.GetGame(clue.GameID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid gameID, categoryID, or clueID"})
		return
	}

	category, err := s.DB.GetCategory(clue.CategoryID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid gameID, categoryID, or clueID"})
		return
	}

	clueJSON := s.jsonHelper(clue)

	cluesForCategory, err := s.DB.GetCluesForCategory(clue.CategoryID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid gameID, categoryID, or clueID"})
		return
	}

	options := []*models.Option{}
	for _, c := range cluesForCategory {
		options = append(options, &models.Option{
			ClueID:   c.ClueID,
			Selected: c.ClueID == clueID,
		})
	}

	debug := struct {
		Clue     *models.Clue
		Game     *models.Game
		Category *models.Category
		Stats    *models.Stats
		Options  []*models.Option
		ClueJSON string
	}{
		Clue:     clue,
		Game:     game,
		Category: category,
		Stats:    s.Stats,
		Options:  options,
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
