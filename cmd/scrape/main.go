package main

import (
	"fmt"

	"github.com/ecshreve/jepp/scraper"
)

func main() {
	gid := int64(6822)

	game, clues := scraper.Scrape(gid)

	db := scraper.NewDB()
	if err := db.InsertGame(&game); err != nil {
		fmt.Println(err)
	}

	for _, clue := range clues {
		if err := db.InsertClue(&clue); err != nil {
			fmt.Println(err)
		}
	}
}
