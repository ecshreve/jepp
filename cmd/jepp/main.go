package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/scraper"
)

func main() {
	db := models.NewDB()
	backfill(db)
	// start := int64(7040)
	// end := int64(8000)

	// for i := start; i < end; i++ {
	// 	if i%10 == 0 {
	// 		utils.SendGotification("jepp", fmt.Sprintf("scraping game %d", i))
	// 	}
	// 	scrapeAndUpdateDB(db, i)
	// 	time.Sleep(1 * time.Second)
	// }
}

func backfill(db *models.JeppDB) {
	log.Fatal("backfilling not maintained")

	batchSize := 100
	iter := 0

	for {
		var cats []*models.Category
		if err := db.Select(&cats, "SELECT * FROM category ORDER BY category_id ASC LIMIT ? OFFSET ?", batchSize, iter*batchSize); err != nil {
			log.Fatal(err)
		}

		if len(cats) == 0 {
			break
		}

		for _, cat := range cats {
			if len(cat.CategoryID) == 16 {
				continue
			}

			cat.CategoryID = models.GetCategoryID(cat.Name)
			if err := db.UpdateCategory(cat); err != nil {
				log.Fatal(err)
			}
		}
		iter += 1
		log.Printf("finished cat batch %d", iter)
	}

	iter = 0
	for {
		var clues []*models.Clue
		if err := db.Select(&clues, "SELECT * FROM clue ORDER BY clue_id ASC LIMIT ? OFFSET ?", batchSize, iter*batchSize); err != nil {
			log.Fatal(err)
		}

		if len(clues) == 0 {
			break
		}

		for _, cl := range clues {
			if len(cl.CategoryID) == 16 {
				continue
			}

			gameCats, _ := db.GetCategoriesForGame(cl.GameID)
			for _, cat := range gameCats {
				if strings.Contains(cat.CategoryID, strings.ToUpper(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(cl.CategoryID, ""))) {
					cl.CategoryID = cat.CategoryID
					break
				}
			}

			if err := db.UpdateClue(cl); err != nil {
				log.Fatal(err)
			}
		}

		iter += 1
		log.Printf("finished clue batch %d", iter)
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
