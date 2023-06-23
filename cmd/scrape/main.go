package main

import (
	"os"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/ecshreve/jepp/pkg/scraper"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("JEPP_LOCAL_DEV") != "true" {
		log.Fatal("this script should only be run in a local development environment")
	}
	log.SetLevel(log.InfoLevel)
	log.Info("Starting Jepp scraper...")

	models.GetDBHandle()

	// Change loop values to scrape different seasons.
	for i := 15; i > 10; i-- {
		if err := scraper.ScrapeSeason(int64(i)); err != nil {
			log.Fatal(err)
		}
	}

	log.Info("...done scraping")
}
