package server

import (
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/pagination"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func (s *Server) registerClueHandlers(rg *gin.RouterGroup) {
	clue := rg.Group("/clues")
	clue.GET("/", pagination.Default(), s.CluesHandler)
	clue.GET("/random", s.RandomClueHandler)
	clue.GET("/:clueID", s.ClueHandler)
}

// CluesHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			game		query	string	false	"Game ID"
//	@Param			category	query	string	false	"Category ID"
//	@Param			page		query	int		false	"Page number"	default(1)
//	@Param			size		query	int		false	"Page size"		default(10)
//	@Success		200			{array}	models.Clue
//	@Router			/clues [get]
func (s *Server) CluesHandler(c *gin.Context) {
	gameID := c.Query("game")
	categoryID := c.Query("category")

	// TODO: fix pagination handling
	var paginationParams models.PaginationParams

	if page, _ := c.Get("page"); page != nil {
		paginationParams.Page = page.(int)
	}

	if size, _ := c.Get("size"); size != nil {
		paginationParams.PageSize = size.(int)
	}

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

	c.JSON(200, clues)
}

// ClueHandler godoc
//
//	@Summary		Returns a clues
//	@Description	Returns a clue
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			clueID	path		string	true	"Clue ID"	default(708002056)
//	@Success		200		{object}	models.Clue
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/clues/{clueID} [get]
func (s *Server) ClueHandler(c *gin.Context) {
	clueID := c.Param("clueID")
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
	clue, err := s.DB.GetRandomClue()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random clue"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clue)
}