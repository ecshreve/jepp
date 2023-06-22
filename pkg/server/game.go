package server

import (
	"net/http"
	"strconv"

	"github.com/ecshreve/jepp/pkg/models"
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
//	@Param			page	query		int	false	"Page number"	default(1)
//	@Param			size	query		int	false	"Page size"		default(10)
//	@Success		200		{array}		models.Game
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/games [get]
func (s *Server) GamesHandler(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("size")
	paginationParams := models.PaginationParams{Page: page, PageSize: size}

	games, err := s.DB.ListGames(&paginationParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, games)
}

// GameHandler godoc
//
//	@Summary		Returns a game
//	@Description	Returns a game

//	@Tags		game
//	@Accept		*/*
//	@Produce	json
//	@Param		gameID	path		string	true	"Game ID"	default(7040)
//	@Success	200		{object}	models.Game
//	@Failure	500		{object}	utils.HTTPError
//	@Router		/games/{gameID} [get]
func (s *Server) GameHandler(c *gin.Context) {
	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to parse gameID %s", gameIDStr))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	game, err := s.DB.GetGame(gameID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get game %d", gameID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, game)
}

// RandomGameHandler godoc
//
//	@Summary		Returns a random game
//	@Description	Returns a random game
//
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{array}		models.Game
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/games/random [get]
func (s *Server) RandomGameHandler(c *gin.Context) {
	game, err := s.DB.GetRandomGame()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random game"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, game)
}
