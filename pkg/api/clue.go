package api

import (
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/pagination"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func (s *Server) registerClueHandlers() {
	s.Router.GET("/clues", pagination.Default(), s.CluesHandler)
	s.Router.GET("/clues/random", s.RandomClueHandler)
	s.Router.GET("/clues/:clueID", s.ClueHandler)
	s.Router.GET("/categories/:categoryID/clues", s.CluesForCategoryHandler)
	s.Router.GET("/games/:gameID/clues", s.CluesForGameHandler)
}

// CluesHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues
//
//	@Tags			clue
//	@Accept			*/*
//	@Produce		json
//	@Param			page	query	int	false	"Page number"	default(1)
//	@Param			size	query	int	false	"Page size"		default(10)
//	@Success		200		{array}	models.Clue
//	@Router			/clues [get]
func (s *Server) CluesHandler(c *gin.Context) {
	page, _ := c.Get("page")
	size, _ := c.Get("size")

	if page == nil || size == nil {
		return
	}

	clues, err := s.DB.ListClues(&models.PaginationParams{Page: page.(int), PageSize: size.(int)})
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

// CluesForCategoryHandler returns a list of clues for a given category.
//
//	@Summary		Returns a list of clues.
//	@Description	Returns a list of clues for a category.
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			categoryID	path		string	true	"Category ID"	default(10LETTERWORDS000)
//	@Success		200			{array}		models.Clue
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/categories/{categoryID}/clues [get]
func (s *Server) CluesForCategoryHandler(c *gin.Context) {
	categoryID := c.Param("categoryID")
	clues, err := s.DB.GetCluesForCategory(categoryID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues for category %s", categoryID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clues)
}

// CluesForGameHandler godoc
//
//	@Summary		Returns a list of clues
//	@Description	Returns a list of clues for a game
//
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Param			gameID	path		string	true	"Game ID"	default(7000)
//	@Success		200		{array}		models.Clue
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/games/{gameID}/clues [get]
func (s *Server) CluesForGameHandler(c *gin.Context) {
	gameID := c.Param("gameID")
	clues, err := s.DB.GetCluesForGame(gameID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues for game %s", gameID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clues)
}
