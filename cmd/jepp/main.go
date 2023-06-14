package main

import (
	"fmt"
	"time"

	"github.com/ecshreve/jepp/pkg/scraper"
)

func main() {
	db := scraper.NewDB()
	start := int64(8045)
	end := int64(8096)

	for i := start; i < end; i++ {
		scrape(db, i)
		time.Sleep(5 * time.Second)
	}
}

func scrape(db *scraper.JeppDB, gid int64) {
	game, clues := scraper.Scrape(gid)

	if err := db.InsertGame(&game); err != nil {
		fmt.Println(err)
	}

	for _, clue := range clues {
		if err := db.InsertClue(&clue); err != nil {
			fmt.Println(err)
		}
	}
}
