package server

import (
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/server/pagination"
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
//	@Param			game		query		string	false	"Game ID"
//	@Param			category	query		string	false	"Category ID"
//	@Param			page		query		int		false	"Page number"	default(1)
//	@Param			size		query		int		false	"Page size"		default(10)
//	@Success		200			{array}		pagination.Response
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/clues [get]
func (s *Server) CluesHandler(c *gin.Context) {
	gameID := c.GetInt64("game")
	categoryID := c.GetInt64("category")
	page := c.GetInt("page")
	size := c.GetInt("limit")
	paginationParams := models.PaginationParams{Page: page, PageSize: size}

	cluesParams := &models.CluesParams{
		GameID:           gameID,
		CategoryID:       categoryID,
		PaginationParams: &paginationParams,
	}

	clues, err := s.DB.GetClues(*cluesParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &pagination.Response{
		Data:  clues,
		Links: pagination.GetLinks(c, int64(len(clues)), &paginationParams),
	})
}

// ClueHandler godoc
//
//	@Summary		Returns a clues
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
	clueID := c.GetInt64("clueID")
	clue, err := s.DB.GetClue(clueID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clue %s", clueID))
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
//	@Tags			clue,random
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
