package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/scraper"
	"github.com/ecshreve/jepp/pkg/utils"
)

func main() {
	db := models.NewDB()
	start := int64(7040)
	end := int64(8000)

	for i := start; i < end; i++ {
		if i%10 == 0 {
			utils.SendGotification("jepp", fmt.Sprintf("scraping game %d", i))
		}
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
