package main

import (
	"github.com/ecshreve/jepp/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := models.NewDB()

	for i := 35; i > 33; i-- {
		log.Infof("scraping season %d ", i)
		gamesForSeason, err := db.GetGamesBySeason(int64(i))
		if err != nil {
			log.Fatal(err)
		}

		cluesForSeason := 0
		for _, game := range gamesForSeason {
			cluesForSeason += scrapeAndUpdateDB(db, game.GameID)
		}
		log.Infof("inserted %d clues and %d games for season %d", cluesForSeason, len(gamesForSeason), i)
	}
}

func scrapeAndUpdateDB(db *models.JeppDB, gid int64) int {
	clues, cats := ScrapeGameClues(gid)

	for clueID, clue := range clues {
		actual, _ := db.GetCategoryByName(cats[clueID])
		if actual != nil {
			clue.CategoryID = actual.CategoryID
			continue
		}

		inserted, err := db.InsertCategory(cats[clueID])
		if err != nil {
			log.Fatal(err)
		}
		clue.CategoryID = inserted.CategoryID
	}

	for _, clue := range clues {
		if err := db.InsertClue(clue); err != nil {
			log.Fatal(err)
		}
	}

	log.Infof("inserted %d clues for game %d", len(clues), gid)
	return len(clues)
}
