package server

import (
	"errors"
	"net/http"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type ClueFilter struct {
	Filter
	GameID     int64 `form:"game"`
	CategoryID int64 `form:"category"`
}

// ClueHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues
//
//	@Tags			api
//	@Accept			*/*
//	@Produce		json
//	@Param			random		query		bool	false	"Random Clue"
//	@Param			id			query		integer	false	"Clue ID"
//	@Param			game		query		integer	false	"Game ID"
//	@Param			category	query		integer	false	"Category ID"
//	@Success		200			{array}		models.Clue
//	@Failure		400			{object}	utils.HTTPError
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/clue [get]
func ClueHandler(c *gin.Context) {
	var filter ClueFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.Error(oops.Wrapf(err, "unable to bind query"))
		utils.NewError(c, http.StatusBadRequest, errors.New("invalid query param"))
		return
	}

	if filter.Random != nil {
		clues, err := mods.GetRandomClueMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random clue"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, clues)
		return
	}

	if filter.ID != nil {
		clue, err := mods.GetClue(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get clue %d", *filter.ID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		cc := []mods.Clue{*clue}
		c.JSON(http.StatusOK, cc)
		return
	}

	cluesParams := &mods.CluesParams{
		GameID:     filter.GameID,
		CategoryID: filter.CategoryID,
		Page:       *filter.Page,
		Limit:      *filter.Limit,
	}

	clues, err := mods.GetClues(*cluesParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, clues)
}
