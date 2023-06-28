package scraper

import (
	"github.com/ecshreve/jepp/pkg/models"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func ScrapeSeason(db *models.JeppDB, i int64) error {
	log.Infof("scraping season %d ", i)
	fillSeasonGames(db, i)

	gamesForSeason, err := db.GetGamesBySeason(int64(i))
	if err != nil {
		return oops.Wrapf(err, "failed to get games for season %d", i)
	}

	cluesForSeason := 0
	for i, game := range gamesForSeason {
		cluesForSeason += ScrapeAndFillCluesForGame(db, game.GameID)
		log.Infof("%d/%d games updated", i, len(gamesForSeason))
	}
	log.Infof("inserted %d clues and %d games for season %d", cluesForSeason, len(gamesForSeason), i)

	return nil
}
