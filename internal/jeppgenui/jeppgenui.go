package jeppui

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ecshreve/jepp/internal/ent"
	"github.com/ecshreve/jepp/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func uiHandler(c *gin.Context) {
	cl, _ := utils.InitDB()

	numClues, err := cl.Clue.Query().Aggregate(ent.Count()).Int(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch number of clues"})
		return
	}

	clue, err := cl.Clue.Query().Offset(rand.Intn(numClues)).First(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch random clue"})
		return
	}

	cat, err := clue.QueryCategory().First(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't fetch category for clue"})
		return
	}

	game, err := clue.QueryGame().First(context.Background())
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, oops.Wrapf(err, "couldn't fetch game for clue"))
		return
	}

	// TODO: validation
	c.HTML(200, "base.html.tpl", gin.H{
		"NumClues": numClues,
		"Clue":     clue,
		"Game":     game,
		"Category": cat,
	})
}

func NewUI(router *gin.Engine) {
	log.SetLevel(log.DebugLevel)

	r := router
	if router == nil {
		r = gin.Default()
	}

	r.StaticFile("/style.css", "./internal/jeppgenui/static/style.css")
	r.StaticFile("/favicon.ico", "./internal/jeppgenui/static/favicon.ico")
	r.StaticFile("/swagger/doc.json", "./internal/ent/openapi.json")
	r.LoadHTMLGlob("internal/jeppgenui/templates/*")
	r.GET("/ui", uiHandler)
}
