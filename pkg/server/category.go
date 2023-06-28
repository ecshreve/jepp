package server

import (
	"net/http"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// CategoryHandler godoc.
//
//	@Summary		Fetch categories
//	@Description	get categories with optional filters
//
//	@Tags			api
//	@Accept			json
//	@Produce		json
//
//	@Param			random	query		bool	false	"If exists or true, returns `limit` random records."
//	@Param			id		query		int64	false	"If exists, returns the record with the given id."
//	@Param			limit	query		int64	false	"Limit the number of records returned."	Default(10)
//
//	@Success		200		{array}		models.Category
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/category [get]
func CategoryHandler(c *gin.Context) {
	var filter Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.Error(oops.Wrapf(err, "unable to bind query"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	if filter.Random != nil {
		category, err := db.GetRandomCategoryMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random category"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, category)
		return
	}

	if filter.ID != nil {
		category, err := db.GetCategory(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get category %d", *filter.ID))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}
		cc := []models.Category{*category}
		c.JSON(http.StatusOK, cc)
		return
	}

	cats, err := db.GetCategories(*filter.Limit)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get categories"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, cats)
}
