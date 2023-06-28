package main

import (
	"encoding/csv"
	"os"

	mods "github.com/ecshreve/jepp/pkg/models"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Starting Jepp dumper...")
	if os.Getenv("JEPP_ENV") != "dev" {
		log.Fatal("dumping should only be done in a dev env...")
	}

	dumpTables()
}

func dumpTables() {
	allSeasons := []mods.Season{}
	allGames := []mods.Game{}
	allCats := []mods.Category{}
	allClues := []mods.Clue{}

	headers := map[string][]string{
		"season":   {"season_id", "start_date", "end_date"},
		"game":     {"game_id", "season_id", "show_num", "game_date", "taped_date"},
		"category": {"category_id", "name"},
		"clue":     {"clue_id", "game_id", "category_id", "question", "answer"},
	}

	f1, _ := os.Create("data/season-dump.csv")
	defer f1.Close()
	f2, _ := os.Create("data/game-dump.csv")
	defer f2.Close()
	f3, _ := os.Create("data/category-dump.csv")
	defer f3.Close()
	f4, _ := os.Create("data/clue-dump.csv")
	defer f4.Close()

	db := sqlx.MustOpen("sqlite3", "data/sqlite/jepp.db")
	if err := db.Select(&allSeasons, "SELECT * FROM season ORDER BY season_id DESC"); err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f1)
	w.Write(headers["season"])
	for _, season := range allSeasons {
		if err := w.Write(season.Dump()); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	if err := db.Select(&allGames, "SELECT * FROM game ORDER BY game_id DESC"); err != nil {
		log.Fatal(err)
	}
	w = csv.NewWriter(f2)
	w.Write(headers["game"])
	for _, game := range allGames {
		if err := w.Write(game.Dump()); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	if err := db.Select(&allCats, "SELECT * FROM category ORDER BY category_id DESC"); err != nil {
		log.Fatal(err)
	}
	w = csv.NewWriter(f3)
	w.Write(headers["category"])
	for _, cat := range allCats {
		if err := w.Write(cat.Dump()); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	if err := db.Select(&allClues, "SELECT * FROM clue ORDER BY clue_id DESC"); err != nil {
		log.Fatal(err)
	}
	w = csv.NewWriter(f4)
	w.Write(headers["clue"])
	for _, clue := range allClues {
		if err := w.Write(clue.Dump()); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	log.Info("done dumping tables to csv...")
}
