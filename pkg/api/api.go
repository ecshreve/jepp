package api

import (
	"net/http"

	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/pagination"
	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server is the API server.
type Server struct {
	ID     string
	Router *gin.Engine
	Clock  clock.Clock
	DB     *models.JeppDB
}

// NewServer returns a new API server.
func NewServer() *Server {
	s := &Server{
		ID:     "SERVER",
		Router: gin.Default(),
		Clock:  clock.New(),
		DB:     models.NewDB(),
	}

	s.registerHandlers()

	return s
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

// CluesForCategoryHandler returns a list of clues for a given category.
//
//	@Summary		Returns a list of clues.
//	@Description	Returns a list of clues for a category.
//
//	@Tags			category
//	@Accept			*/*
//	@Produce		json
//	@Param			categoryID	path		string	true	"Category ID"	default(10LETTERWORDS000)
//	@Success		200			{array}		models.Clue
//	@Failure		500			{object}	utils.HTTPError
//	@Router			/categories/{categoryID}/clues [get]
func (s *Server) CluesForCategoryHandler(c *gin.Context) {
	categoryID := c.Param("categoryID")
	clues, err := s.DB.GetCluesForCategory(categoryID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues for category %s", categoryID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clues)
}

// CluesForGameHandler godoc
//
//	@Summary		Returns a list of clues.
//	@Description	Returns a list of clues for a game.

// @Tags			game
// @Accept			*/*
// @Produce		json
// @Param			gameID	path		string	true	"Game ID"	default(7000)
// @Success		200		{array}		models.Clue
// @Failure		500		{object}	utils.HTTPError
// @Router			/games/{gameID}/clues [get]
func (s *Server) CluesForGameHandler(c *gin.Context) {
	gameID := c.Param("gameID")
	clues, err := s.DB.GetCluesForGame(gameID)
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get clues for game %s", gameID))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, clues)
}

// GamesHandler godoc
//
//	@Summary		Returns a list of games
//	@Description	Returns a list of games
//	@Tags			game
//	@Accept			*/*
//	@Produce		json
//	@Param			page	query	int	false	"Page number"	default(1)
//	@Param			size	query	int	false	"Page size"		default(10)
//	@Success		200		{array}	models.Game
//	@Router			/games [get]
func (s *Server) GamesHandler(c *gin.Context) {
	page, _ := c.Get("page")
	size, _ := c.Get("size")

	if page == nil || size == nil {
		return
	}

	games, err := s.DB.ListGames(&models.PaginationParams{Page: page.(int), PageSize: size.(int)})
	if err != nil {
		log.Error(oops.Wrapf(err, "unable to get games"))
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(200, games)
}

// BaseHandler godoc
//
//	@Summary		Base handler
//	@Description	Show available endpoints

// @Tags			root
// @Accept			*/*
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/ [get]
func (s *Server) BaseHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api root",
		"available endpoints": []string{
			"/",
			"/ping",
			"/games",
			"/games/:gameID/clues",
			"/categories",
			"/categories/:categoryID/clues",
		},
	})
}

// PingHandler godoc
//
//	@Summary		Show the status of server
//	@Description	Get the status of server

// @Tags			root
// @Accept			*/*
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Router			/ping [get]
func (s *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// registerHandlers registers all the handlers for the server.
func (s *Server) registerHandlers() {
	s.Router.GET("/", s.BaseHandler)
	s.Router.GET("/ping", s.PingHandler)
	s.Router.GET("/games", pagination.Default(), s.GamesHandler)
	s.Router.GET("/games/:gameID/clues", s.CluesForGameHandler)
	s.Router.GET("/categories", pagination.Default(), s.CategoriesHandler)
	s.Router.GET("/categories/:categoryID/clues", s.CluesForCategoryHandler)
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// This is not safe for production use but it's fine for playing
	// around locally.
	if err := s.Router.SetTrustedProxies(nil); err != nil {
		log.Error(oops.Wrapf(err, "unable to set proxies"))
	}
}

// Serve starts the server.
func (s *Server) Serve() error {
	err := s.Router.Run(":8880")
	if err != nil {
		return oops.Wrapf(err, "gin server returned error")
	}
	return nil
}
