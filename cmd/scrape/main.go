package main

import (
	"log"
	"time"

	"github.com/ecshreve/jepp/pkg/models"
)

func main() {
	db := models.NewDB()
	start := int64(7050)
	end := int64(8000)

	// var cnt int64
	for i := start; i < end; i++ {
		// if err := db.Get(&cnt, "SELECT COUNT(DISTINCT game_id) FROM category"); err != nil {
		// 	log.Fatal(err)
		// }
		// if cnt > 10 {
		// 	utils.SendGotification("jepp", "done")
		// 	return
		// }

		// if i%10 == 0 {
		// 	utils.SendGotification("jepp", fmt.Sprintf("scraping game %d", i))
		// }
		scrapeAndUpdateDB(db, i)
		time.Sleep(1 * time.Second)
	}
}

func scrapeAndUpdateDB(db *models.JeppDB, gid int64) {
	game, clues, cats := Scrape(gid)

	if err := db.InsertGame(&game); err != nil {
		log.Fatal(err)
	}

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
}
