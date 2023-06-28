package server

import (
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// GameHandler godoc
//
//	@Summary		Fetch Games
//	@Description	get games with optional filters
//
//	@Tags			api
//	@Accept			json
//	@Produce		json
//
//	@Param			random	query		bool	false	"If exists or true, returns `limit` random records."
//	@Param			id		query		int64	false	"If exists, returns the record with the given id."
//	@Param			limit	query		int64	false	"Limit the number of records returned."	Default(10)
//
//	@Success		200		{array}		models.Game
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/game [get]
func GameHandler(c *gin.Context) {
	var filter Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.Error(oops.Wrapf(err, "unable to bind query"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	if filter.Random != nil {
		games, err := models.GetRandomGameMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random game"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, games)
		return
	}

	if filter.ID != nil {
		game, err := models.GetGame(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get game %d", *filter.ID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		gg := []models.Game{*game}
		c.JSON(http.StatusOK, gg)
		return
	}

	games, err := models.GetGames(*filter.Limit)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, games)
}
