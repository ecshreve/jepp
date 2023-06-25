package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StringResponse struct {
	Message string `json:"message" example:"some string value"`
}

// BaseAPIHandler handles the base API route.
func BaseAPIHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []string{
		"/",
		"/games",
		"/categories",
		"/clues",
	})
}
