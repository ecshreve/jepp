package server

import (
	"net/http"
	"strconv"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/pagination"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func (s *Server) registerGameHandlers(rg *gin.RouterGroup) {
	games := rg.Group("/games")
	games.GET("/", pagination.Default(), s.GamesHandler)
	games.GET("/random", s.RandomGameHandler)
	games.GET("/:gameID", s.GameHandler)
}

// GamesHandler godoc
//
//	@Summary		Returns a list of games
//	@Description	Returns a list of games
//
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Param			page	query	int	false	"Page number"	default(1)
//	@Param			size	query	int	false	"Page size"		default(10)
//	@Success		200		{array}	models.Game
//	@Router			/games [get]
func (s *Server) GamesHandler(c *gin.Context) {
	page, _ := c.Get("page")
	size, _ := c.Get("size")

	if page == nil || size == nil {
		return
	}

	games, err := s.DB.ListGames(&models.PaginationParams{Page: page.(int), PageSize: size.(int)})
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, games)
}

// GameHandler godoc
//
//	@Summary		Returns a games
//	@Description	Returns a game
//
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Param			gameID	path		string	true	"Game ID"	default(7000)
//	@Success		200		{object}	models.Game
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/games/{gameID} [get]
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
		log.Error(oops.Wrapf(err, "unable to get game %s", gameID))
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
//	@Tags			game,random
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
