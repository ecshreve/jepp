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

func (s *Server) registerCategoryHandlers() {
	s.Router.GET("/categories", pagination.Default(), s.CategoriesHandler)
	s.Router.GET("/categories/random", s.RandomCategoryHandler)
	s.Router.GET("/categories/:categoryID", s.CategoryHandler)
}

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
//	@Success		200		{array}		models.CategoryCount
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/categories [get]
func (s *Server) CategoriesHandler(c *gin.Context) {
	page, _ := c.Get("page")
	size, _ := c.Get("size")

	if page == nil || size == nil {
		return
	}

	cats, err := s.DB.ListCategories(&models.PaginationParams{Page: page.(int), PageSize: size.(int)})
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get categories"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, cats)
}

// CategoryHandler godoc
//
//	@Summary		Returns a category
//	@Description	Returns a category
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			categoryID	path		string	true	"Category ID"	default(10LETTERWORDS000)
//	@Success		200			{object}	models.CategoryCount
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/categories/{categoryID} [get]
func (s *Server) CategoryHandler(c *gin.Context) {
	categoryID := c.Param("categoryID")
	category, err := s.DB.GetCategory(categoryID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get category %s", categoryID))
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
//	@Success		200	{object}	models.CategoryCount
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
