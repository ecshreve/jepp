package main

import (
	"fmt"
	"time"

	"github.com/ecshreve/jepp/scraper"
)

func main() {
	start := int64(6916)
	end := int64(8000)

	db := scraper.NewDB()
	for i := start; i < end; i++ {
		time.Sleep(1 * time.Second)
		game, clues := scraper.Scrape(i)

		if err := db.InsertGame(&game); err != nil {
			fmt.Println(err)
		}

		for _, clue := range clues {
			if err := db.InsertClue(&clue); err != nil {
				fmt.Println(err)
			}
		}
	}
}
