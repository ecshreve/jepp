package server

import (
	"errors"
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
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
//	@Summary		Fetch Clues
//	@Description	get clues with optional filters
//
//	@Tags			api
//	@Accept			json
//	@Produce		json
//
//	@Param			random		query		bool	false	"If exists or true, returns `limit` random records."
//	@Param			id			query		int64	false	"If exists, returns the record with the given id."
//	@Param			game		query		integer	false	"Filter by Game ID"
//	@Param			category	query		integer	false	"Filter by Category ID"
//	@Param			limit		query		int64	false	"Limit the number of records returned"	Default(10)
//
//	@Success		200			{array}		models.Clue
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
		clues, err := db.GetRandomClueMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random clue"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, clues)
		return
	}

	if filter.ID != nil {
		clue, err := db.GetClue(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get clue %d", *filter.ID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		cc := []models.Clue{*clue}
		c.JSON(http.StatusOK, cc)
		return
	}

	cluesParams := &models.CluesParams{
		GameID:     filter.GameID,
		CategoryID: filter.CategoryID,
		Page:       *filter.Page,
		Limit:      *filter.Limit,
	}

	clues, err := db.GetClues(*cluesParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, clues)
}
