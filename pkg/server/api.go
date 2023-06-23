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
//	@Tags			api
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{array}	string	"list of available endpoints"
//	@Router			/ [get]
func BaseHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []string{
		"/api",
		"/api/ping",
		"/api/games",
		"/api/categories",
		"/api/clues",
		"/api/random/clue",
		"/api/random/category",
		"/api/random/game",
	})
}

// PingHandler godoc
//
//	@Summary		Show the status of server
//	@Description	Get the status of server
//
//	@Tags			api
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{string}	pong
//	@Router			/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
