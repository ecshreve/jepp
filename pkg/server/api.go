package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseHandler godoc
//
//	@Summary		Base api handler
//	@Description	List available endpoints
//
//	@Tags			api
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/api [get]
func BaseHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"available endpoints": []string{
			"/api",
			"/api/ping",
			"/api/games",
			"/api/categories",
			"/api/clues",
			"/api/random/clue",
			"/api/random/category",
			"/api/random/game",
		},
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
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
