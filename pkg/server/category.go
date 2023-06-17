package server

import (
	"net/http"
	"strconv"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/server/pagination"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// CategoriesHandler returns a list of categories.
//
//	@Summary		Returns a list of categories.
//	@Description	Returns a list of categories.
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			page	query		int	false	"Page number"	default(1)
//	@Param			size	query		int	false	"Page size"		default(10)
//	@Success		200		{object}	pagination.Response
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/categories [get]
func (s *Server) CategoriesHandler(c *gin.Context) {
	page := c.GetInt("page")
	size := c.GetInt("limit")
	paginationParams := models.PaginationParams{Page: page, PageSize: size}

	cats, err := s.DB.GetCategories(paginationParams)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get categories"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &pagination.Response{
		Data:  cats,
		Links: pagination.GetLinks(c, int64(len(cats)), &paginationParams),
	})
}

// CategoryHandler godoc
//
//	@Summary		Returns a category
//	@Description	Returns a category
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			categoryID	path		int	true	"Category ID"	default(10LETTERWORDS000)
//	@Success		200			{object}	models.Category
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/categories/{categoryID} [get]
func (s *Server) CategoryHandler(c *gin.Context) {
	categoryIDStr := c.Param("categoryID")
	categoryID, _ := strconv.ParseInt(categoryIDStr, 10, 64)
	category, err := s.DB.GetCategory(categoryID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get category %d", categoryID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, category)
}

// RandomCategoryHandler godoc
//
//	@Summary		Returns a random category
//	@Description	Returns a random category
//
//	@Tags			category,random
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	models.Category
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/categories/random [get]
func (s *Server) RandomCategoryHandler(c *gin.Context) {
	category, err := s.DB.GetRandomCategory()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random category"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, category)
}
