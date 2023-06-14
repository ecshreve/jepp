package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// BBallDB is a wrapper around sqlx.DB.
type JeppDB struct {
	*sqlx.DB
}

func NewDB() *JeppDB {
	dbname := "jeppdb"
	addr := fmt.Sprintf("%s:3306", os.Getenv("DB_HOST"))

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "jepp-user",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Get a database handle.
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	jdb := &JeppDB{db}
	jdb.initDB()
	return jdb
}

func (db *JeppDB) initDB() {
	game := `
CREATE TABLE IF NOT EXISTS game (
    game_id INT NOT NULL PRIMARY KEY,
    show_num INT NOT NULL,
    game_date DATE NOT NULL
);`

	clue := `
CREATE TABLE IF NOT EXISTS clue (
    clue_id VARCHAR(64) NOT NULL,
    game_id INT NOT NULL,
    category VARCHAR(256) NOT NULL,
    question TEXT NOT NULL,
    answer VARCHAR(256) NOT NULL,
    PRIMARY KEY (clue_id, game_id),
		CONSTRAINT fk_clue_game FOREIGN KEY (game_id) REFERENCES game(game_id)
);`

	db.MustExec(game)
	db.MustExec(clue)
}
