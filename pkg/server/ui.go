package server

import (
	"encoding/json"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
)

// BaseUIHandler handles the base UI route.
func BaseUIHandler(c *gin.Context) {
	clue, err := mods.GetRandomClue()
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch random clue"})
		return
	}

	cat, err := mods.GetCategory(clue.CategoryID)
	if err != nil {
		c.JSON(400, gin.H{"error": "couldn't fetch category for clue"})
		return
	}

	game, err := mods.GetGame(clue.GameID)
	if err != nil {
		utils.NewError(c, 400, oops.Wrapf(err, "couldn't fetch game for clue"))
		return
	}

	clueJSON, _ := json.Marshal(clue)
	numClues, _ := mods.NumClues()

	c.HTML(200, "base.html.tpl", gin.H{
		"NumClues": numClues,
		"Clue":     clue,
		"ClueJSON": string(clueJSON),
		"Game":     game,
		"Category": cat,
	})
}
