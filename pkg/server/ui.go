package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// BaseUIHandler handles the base UI route.
func (s *Server) BaseUIHandler(c *gin.Context) {
	clue, err := s.DB.GetRandomClue()
	if err != nil {
		log.Error(oops.Wrapf(err, "couldn't fetch random clue"))
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
