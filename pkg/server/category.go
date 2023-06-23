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

// CategoriesHandler returns a list of categories.
//
//	@Summary		Returns a list of categories.
//	@Description	Returns a list of categories.
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			id	query		int64	false	"Category ID"
//	@Success		200	{array}		models.Category
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/categories [get]
func CategoriesHandler(c *gin.Context) {
	categoryIDStr := c.DefaultQuery("id", "")
	if categoryIDStr != "" {
		categoryID, _ := strconv.ParseInt(categoryIDStr, 10, 64)
		category, err := mods.GetCategory(categoryID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get category %d", categoryID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}
		cc := []mods.Category{*category}
		c.JSON(http.StatusOK, cc)
		return
	}

	cats, err := mods.GetCategories()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get categories"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, cats)
}

// RandomCategoryHandler godoc
//
//	@Summary		Returns a random category
//	@Description	Returns a random category
//
//	@Tags			random
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	models.Category
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/random/category [get]
func RandomCategoryHandler(c *gin.Context) {
	category, err := mods.GetRandomCategory()
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get random category"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, category)
}
