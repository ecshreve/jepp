package main

import (
	"github.com/ecshreve/jepp/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := models.NewDB()

	gamesForSeason, err := db.GetGamesBySeason(37)
	if err != nil {
		log.Fatal(err)
	}

	for _, game := range gamesForSeason {
		scrapeAndUpdateDB(db, game.GameID)
	}
}

func scrapeAndUpdateDB(db *models.JeppDB, gid int64) {
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
}
