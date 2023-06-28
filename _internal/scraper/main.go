package scraper

import (
	"os"

	"github.com/ecshreve/jepp/pkg/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("JEPP_ENV") != "dev" {
		log.Fatal("this script should only be run in a local development environment")
	}
	log.SetLevel(log.DebugLevel)
	log.Info("Starting Jepp scraper...")

	// jdb := models.NewJeppDB("data/jepp.db")
	jdb := models.NewJeppDB()
	num := ScrapeAndFillCluesForGame(jdb, 3109)
	log.Infof("inserted %d clues", num)
	// Change loop values to scrape different seasons.
	// for i := 25; i > 24; i-- {
	// 	if err := scraper.ScrapeSeason(jdb, int64(i)); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	log.Info("...done scraping")
}
