package server

import (
	"encoding/json"
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
)

// BaseUIHandler handles the base UI route.
func BaseUIHandler(c *gin.Context) {
	clue, err := models.GetRandomClue()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch random clue"})
		return
	}

	cat, err := models.GetCategory(clue.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch category for clue"})
		return
	}

	game, err := models.GetGame(clue.GameID)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, oops.Wrapf(err, "couldn't fetch game for clue"))
		return
	}

	// TODO: validation
	clueJSON, _ := json.Marshal(clue)
	numClues, _ := models.CountClues()

	c.HTML(200, "base.html.tpl", gin.H{
		"NumClues": numClues,
		"Clue":     clue,
		"ClueJSON": string(clueJSON), // FIX: this is ugly
		"Game":     game,
		"Category": cat,
	})
}
