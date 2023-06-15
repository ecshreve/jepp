package main

import (
	"log"
	"time"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/scraper"
)

func main() {
	db := models.NewDB()
	start := int64(8000)
	end := int64(8500)

	for i := start; i < end; i++ {
		scrapeAndUpdateDB(db, i)
		time.Sleep(1 * time.Second)
	}
}

func scrapeAndUpdateDB(db *models.JeppDB, gid int64) {
	game, clues, cats := scraper.Scrape(gid)

	if err := db.InsertGame(&game); err != nil {
		log.Fatal(err)
	}

	for _, cat := range cats {
		if err := db.InsertCategory(&cat); err != nil {
			log.Fatal(err)
		}
	}

	for _, clue := range clues {
		if err := db.InsertClue(&clue); err != nil {
			log.Fatal(err)
		}
	}
}
