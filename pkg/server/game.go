package server

import (
	"net/http"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// GameHandler godoc
//
//	@Summary		Returns a list of games
//	@Description	Returns a list of games
//
//	@Tags			api
//	@Accept			*/*
//	@Produce		json
//	@Param			filter	query		Filter	false	"Filter games"
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
		games, err := mods.GetRandomGameMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random game"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, games)
		return
	}

	if filter.ID != nil {
		game, err := mods.GetGame(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get game %d", *filter.ID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		gg := []mods.Game{*game}
		c.JSON(http.StatusOK, gg)
		return
	}

	games, err := mods.GetGames()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, games)
}
