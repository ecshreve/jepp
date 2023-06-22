package server

import (
	"fmt"
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
	// s.Router.POST("/", s.BaseUIHandler)
	s.Router.GET("/:clueID", s.ClueUIHandler)
	s.Router.POST("/:clueID", s.ClueUIPOSTHandler)
	s.Router.GET("/quiz", s.QuizHandler)
	s.Router.POST("/quiz", s.QuizHandler)

	s.Router.GET("/debug", s.DebugUIHandler)

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

	c.HTML(200, "landing.html.tpl", gin.H{
		"Stats":    s.Stats,
		"Clue":     clue,
		"Game":     game,
		"Category": cat,
	})
}

func (s *Server) QuizHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		cor := c.PostForm("correct")
		inc := c.PostForm("incorrect")

		correct := len(cor) > len(inc)
		if correct {
			s.QZ.Correct++
		} else {
			s.QZ.Incorrect++
		}
		s.QZ.Total++
	}

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

	s.QZ.Clues = append(s.QZ.Clues, clue)

	c.HTML(200, "quiz.html.tpl", gin.H{
		"Clue":     clue,
		"Category": cat,
		"Session":  s.QZ,
		"Game":     game,
		"Viz":      Viz(*s.QZ),
	})
}

// ClueUIPOSTHandler godoc
func (s *Server) ClueUIPOSTHandler(c *gin.Context) {
	clueIDStr := c.Param("clueID")
	clueID, err := strconv.ParseInt(clueIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid clueID"})
		return
	}

	clue, err := s.DB.GetClue(clueID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch clue for clueID"})
		return
	}

	categoryIDStr := c.PostForm("cat-sel")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid categoryID"})
		return
	}

	clues, err := s.DB.ListClues(models.CluesParams{GameID: clue.GameID, CategoryID: categoryID})
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch clues for categoryID"})
		return
	}
	log.Infof("clues: %s", pretty.Sprint(clues))

	newClueIDStr := fmt.Sprintf("%d", clues[0].ClueID)
	c.Params = gin.Params{gin.Param{Key: "clueIDStr", Value: newClueIDStr}}
	c.Request.Method = "GET"
	c.Redirect(302, "/"+newClueIDStr)
}

// ClueUIHandler godoc
func (s *Server) ClueUIHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		s.ClueUIPOSTHandler(c)
		return
	}

	clueID, err := strconv.ParseInt(c.Param("clueID"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid clueID"})
		return
	}

	clue, err := s.DB.GetClue(clueID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch clue for clueID"})
		return
	}

	game, err := s.DB.GetGame(clue.GameID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch game for clue"})
		return
	}

	category, err := s.DB.GetCategory(clue.CategoryID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch category for clue"})
		return
	}

	clueJSON := s.jsonHelper(clue)

	categoriesForGame, err := s.DB.GetCategoriesForGame(clue.GameID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch categories for game"})
		return
	}

	catOpts := []models.Option{}
	for _, c := range categoriesForGame {
		catOpts = append(catOpts, models.Option{
			OptionKey: fmt.Sprintf("%d", c.CategoryID),
			OptionVal: c.Name,
			Selected:  c.CategoryID == clue.CategoryID,
		})
	}

	cluesForGame, err := s.DB.ListClues(models.CluesParams{GameID: clue.GameID})
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch clues for game"})
		return
	}

	nextClue := models.Clue{}
	prevClue := models.Clue{}
	for i, c := range cluesForGame {
		if c.ClueID == clue.ClueID {
			if i > 0 {
				prevClue = *cluesForGame[i+1]
			}
			if i < len(cluesForGame)-1 {
				nextClue = *cluesForGame[i-1]
			}
		}
	}

	navLinks := models.NavLinks{
		NextClue: fmt.Sprintf("%d", nextClue.ClueID),
		PrevClue: fmt.Sprintf("%d", prevClue.ClueID),
	}

	options := &models.Options{
		ClueID:          clue.ClueID,
		Links:           navLinks,
		CategoryOptions: catOpts,
	}

	debug := struct {
		Clue     *models.Clue
		Game     *models.Game
		Category *models.Category
		Stats    *models.Stats
		Options  *models.Options
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
