package server

import (
	"github.com/gin-gonic/gin"
)

// BaseHandler godoc
//
//	@Summary		Base handler
//	@Description	Show available endpoints
//
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ [get]
func BaseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api root",
		"available endpoints": []string{
			"/api",
			"/api/ping",
			"/api/games",
			"/api/games/random",
			"/api/games/:gameID",
			"/api/categories",
			"/api/categories/random",
			"/api/categories/:categoryID",
			"/api/clues",
			"/api/clues/random",
			"/api/clues/:clueID",
		},
	})
}

// PingHandler godoc
//
//	@Summary		Show the status of server
//	@Description	Get the status of server
//
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
