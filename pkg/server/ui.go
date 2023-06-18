package server

import (
	"strconv"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// registerUIHandlers registers route handlers for the UI.
func (s *Server) registerUIHandlers() {
	s.Router.StaticFile("style.css", "./static/style.css")
	s.Router.StaticFile("favicon.ico", "./static/favicon.ico")

	s.Router.LoadHTMLGlob("pkg/server/templates/*")

	s.Router.GET("/", s.BaseUIHandler)
	s.Router.POST("/", s.BaseUIHandler)

	s.Router.GET("/debug", s.DebugUIHandler)

	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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

func (s *Server) BaseUIHandler(c *gin.Context) {
	if clueIDStr := c.PostForm("json-button"); clueIDStr != "" {
		clueID, _ := strconv.ParseInt(clueIDStr, 10, 64)

		clue, _ := s.DB.GetClue(clueID)
		cat, _ := s.DB.GetCategory(clue.CategoryID)
		game, _ := s.DB.GetGame(clue.GameID)

		c.JSON(200, gin.H{
			"clueID":       clue.ClueID,
			"categoryID":   clue.CategoryID,
			"categoryName": cat.Name,
			"gameID":       clue.GameID,
			"gameDate":     game.GameDate,
			"question":     clue.Question,
			"answer":       clue.Answer,
		})
		return
	}

	var clues []*models.Clue
	rolls := []string{"game-roll", "cat-roll", "clue-roll"}
	for _, roll := range rolls {
		if clueID := c.PostForm(roll); clueID != "" {
			clues, _ = s.PostFormHandler(roll, clueID)
			log.Info(roll)
		}
	}

	clue, _ := s.DB.GetRandomClue(clues)
	game, _ := s.DB.GetGame(clue.GameID)
	category, _ := s.DB.GetCategory(clue.CategoryID)
	gc, _ := s.DB.GetCategoryGameCount(clue.CategoryID)
	cc, _ := s.DB.GetCategoryClueCount(clue.CategoryID)

	random := struct {
		*models.Clue
		*models.Game
		*models.Category
		*models.Stats
		CategoryGamesCount int64
		CategoryCluesCount int64
		ClueJSON           string
	}{
		Clue:               clue,
		Game:               game,
		Category:           category,
		Stats:              s.Stats,
		CategoryGamesCount: gc,
		CategoryCluesCount: cc,
		ClueJSON: pretty.Sprint(gin.H{
			"clueID":       clue.ClueID,
			"gameID":       clue.GameID,
			"gameDate":     game.GameDate,
			"categoryID":   clue.CategoryID,
			"categoryName": category.Name,
			"question":     clue.Question,
			"answer":       clue.Answer,
		}),
	}
	c.HTML(200, "base.html.tpl", random)
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
