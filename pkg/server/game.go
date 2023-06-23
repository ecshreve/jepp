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

// GamesHandler godoc
//
//	@Summary		Returns a list of games
//	@Description	Returns a list of games
//
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Param			id	query		int64	false	"Game ID"
//	@Success		200	{array}		models.Game
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/games [get]
func GamesHandler(c *gin.Context) {
	gameIDStr := c.DefaultQuery("id", "")
	if gameIDStr != "" {
		gameID, _ := strconv.ParseInt(gameIDStr, 10, 64)
		game, err := mods.GetGame(gameID)
		if err != nil || game == nil {
			log.Error(oops.Wrapf(err, "unable to get game %d", gameID))
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

// RandomGameHandler godoc
//
//	@Summary		Returns a random game
//	@Description	Returns a random game
//
//	@Tags			random
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	models.Game
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/random/game [get]
func RandomGameHandler(c *gin.Context) {
	game, err := mods.GetRandomGame()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random game"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, game)
}
