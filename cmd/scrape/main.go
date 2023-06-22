package main

import (
	"fmt"
	"os"

	"github.com/ecshreve/jepp/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("JEPP_LOCAL_DEV") != "true" {
		log.Fatal("this script should only be run in a local development environment")
	}

	dbname := "jeppdb"
	dbuser := "jepp-user"
	dbpass := os.Getenv("MYSQL_USER_PASS")
	dbaddr := fmt.Sprintf("%s:3306", os.Getenv("DB_HOST"))
	db := models.NewDB(dbname, dbuser, dbpass, dbaddr)

	// Change loop values to scrape different seasons.
	for i := 38; i > 38; i-- {
		log.Infof("scraping season %d ", i)
		gamesForSeason, err := db.GetGamesBySeason(int64(i))
		if err != nil {
			log.Fatal(err)
		}

		cluesForSeason := 0
		for i, game := range gamesForSeason {
			cluesForSeason += scrapeAndFillCluesForGame(db, game.GameID)
			log.Infof("%d/%d games updated", i, len(gamesForSeason))
		}
		log.Infof("inserted %d clues and %d games for season %d", cluesForSeason, len(gamesForSeason), i)
	}
}
