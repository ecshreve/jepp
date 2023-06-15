package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	return jdb
}

// InitDB initializes the database.
//
// Not currently used. I manually create tables in Adminer.
func (db *JeppDB) initDB() {
	db.MustExec(GAME_SCHEMA)
	db.MustExec(CLUE_SCHEMA)
	db.MustExec(CATEGORY_SCHEMA)
}

func (db *JeppDB) Dump() error {
	games, _ := db.GetAllGames()
	fmt.Println(games[0])
	file, _ := json.MarshalIndent(games, "", " ")
	_ = ioutil.WriteFile("game-dump.json", file, 0644)

	clues, _ := db.GetAllClues()
	fmt.Println(clues[0])
	file, _ = json.MarshalIndent(clues, "", " ")
	_ = ioutil.WriteFile("clue-dump.json", file, 0644)
	return nil
}
