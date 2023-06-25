package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StringResponse struct {
	Message string `json:"message" example:"some string value"`
}

// BaseHandler godoc
//
//	@Summary		Base api handler
//	@Description	List available endpoints
//
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{array}	string	"list of available endpoints"
//	@Router			/ [get]
func BaseAPIHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []string{
		"/",
		"/games",
		"/categories",
		"/clues",
	})
}
