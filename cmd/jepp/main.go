package main

import (
	"fmt"
	"time"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/scraper"
)

func main() {
	db := models.NewDB()
	start := int64(8000)
	end := int64(8044)

	for i := start; i < end; i++ {
		scrape(db, i)
		time.Sleep(1 * time.Second)
	}
}

func scrape(db *models.JeppDB, gid int64) {
	game, clues, cats := scraper.Scrape(gid)

	if err := db.InsertGame(&game); err != nil {
		fmt.Println(err)
	}

	for _, clue := range clues {
		if err := db.InsertClue(&clue); err != nil {
			fmt.Println(err)
		}
	}

	for _, cat := range cats {
		if err := db.InsertCategory(&cat); err != nil {
			fmt.Println(err)
		}
	}
}
