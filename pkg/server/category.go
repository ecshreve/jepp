package server

import (
	"net/http"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// CategoryHandler godoc.
//
//	@Summary		Returns a list of categories.
//	@Description	Returns a list of categories.
//	@Tags			api
//	@Accept			json
//	@Produce		json
//	@Param			random	query	bool	false	"If exists, returns up to `limit` random records."
//	@Param			id		query	int64	false	"If exists, returns the record with the given id."
//	@Param			page	query	int64	false	"Paging offset"
//	@Param			limit	query	int64	false	"Limit the number of records returned"
//	@Success		200		{array}	models.Category
//	@Router			/category [get]
func CategoryHandler(c *gin.Context) {
	var filter Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		log.Error(oops.Wrapf(err, "unable to bind query"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}
	log.Debugf("filter: %#v", filter)

	if filter.Random != nil {
		category, err := mods.GetRandomCategoryMany(*filter.Limit)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get random category"))
			utils.NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, category)
		return
	}

	if filter.ID != nil {
		category, err := mods.GetCategory(*filter.ID)
		if err != nil {
			log.Error(oops.Wrapf(err, "unable to get category %d", *filter.ID))
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
