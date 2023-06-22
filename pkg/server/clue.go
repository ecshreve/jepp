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

// CluesHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			game		query		int64	false	"Game ID"
//	@Param			category	query		int64	false	"Category ID"
//	@Param			page		query		int		false	"Page number"	default(1)
//	@Param			size		query		int		false	"Page size"		default(10)
//	@Success		200			{array}		models.Clue
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/clues [get]
func (s *Server) CluesHandler(c *gin.Context) {
	gameIDStr := c.Query("game")
	gameID, _ := strconv.ParseInt(gameIDStr, 10, 64)

	categoryIDStr := c.Query("category")
	categoryID, _ := strconv.ParseInt(categoryIDStr, 10, 64)

	page := c.GetInt("page")
	size := c.GetInt("size")
	paginationParams := models.PaginationParams{Page: page, PageSize: size}

	cluesParams := &models.CluesParams{
		GameID:           gameID,
		CategoryID:       categoryID,
		PaginationParams: &paginationParams,
	}

	clues, err := s.DB.ListClues(*cluesParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, clues)
}

// ClueHandler godoc
//
//	@Summary		Returns a clue
//	@Description	Returns a clue
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			clueID	path		int64	true	"Clue ID"	default(708002056)
//	@Success		200		{object}	models.Clue
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/clues/{clueID} [get]
func (s *Server) ClueHandler(c *gin.Context) {
	clueIDStr := c.Param("clueID")
	clueID, _ := strconv.ParseInt(clueIDStr, 10, 64)
	clue, err := s.DB.GetClue(clueID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clue %d", clueID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clue)
}

// RandomClueHandler godoc
//
//	@Summary		Returns a random clue
//	@Description	Returns a random clue
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{array}		models.Clue
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/clues/random [get]
func (s *Server) RandomClueHandler(c *gin.Context) {
	clue, err := s.DB.GetRandomClue(nil)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random clue"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clue)
}
