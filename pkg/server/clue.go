package server

import (
	"net/http"
	"strconv"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// CluesHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			id			query		integer	false	"Clue ID"
//	@Param			game		query		integer	false	"Game ID"
//	@Param			category	query		integer	false	"Category ID"
//	@Success		200			{array}		models.Clue
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/clues [get]
func CluesHandler(c *gin.Context) {
	clueIDStr := c.DefaultQuery("id", "")
	if clueIDStr != "" {
		clueID, _ := strconv.ParseInt(clueIDStr, 10, 64)
		clue, err := mods.GetClue(clueID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get clue %d", clueID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		cc := []mods.Clue{*clue}
		c.JSON(http.StatusOK, cc)
		return
	}

	gameIDStr := c.Query("game")
	gameID, _ := strconv.ParseInt(gameIDStr, 10, 64)

	categoryIDStr := c.Query("category")
	categoryID, _ := strconv.ParseInt(categoryIDStr, 10, 64)

	cluesParams := &mods.CluesParams{
		GameID:     gameID,
		CategoryID: categoryID,
	}

	clues, err := mods.GetClues(*cluesParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, clues)
}

// RandomClueHandler godoc
//
//	@Summary		Returns a random clue
//	@Description	Returns a random clue
//
//	@Tags			random
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	models.Clue
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/random/clue [get]
func RandomClueHandler(c *gin.Context) {
	clue, err := mods.GetRandomClue()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random clue"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, clue)
}
