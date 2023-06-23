package main

import (
	"os"

	"github.com/ecshreve/jepp/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("JEPP_LOCAL_DEV") != "true" {
		log.Fatal("this script should only be run in a local development environment")
	}

	// Change loop values to scrape different seasons.
	for i := 38; i > 38; i-- {
		log.Infof("scraping season %d ", i)
		gamesForSeason, err := models.GetGamesBySeason(int64(i))
		if err != nil {
			log.Fatal(err)
		}

		cluesForSeason := 0
		for i, game := range gamesForSeason {
			cluesForSeason += scrapeAndFillCluesForGame(nil, game.GameID)
			log.Infof("%d/%d games updated", i, len(gamesForSeason))
		}
		log.Infof("inserted %d clues and %d games for season %d", cluesForSeason, len(gamesForSeason), i)
	}
}
