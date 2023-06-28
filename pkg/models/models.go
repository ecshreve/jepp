package models

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type JeppDB struct {
	*sqlx.DB
}

var db *sqlx.DB

func NewJeppDB() *JeppDB {
	dbfile := "data/sqlite/jepp.db"
	db = sqlx.MustOpen("sqlite3", dbfile)
	log.Debugf("NEW -- %v", db != nil)
	log.Debug(os.Getenv("PWD"))

	return &JeppDB{DB: db}
}

func NewTestJeppDB() *JeppDB {
	db = sqlx.MustOpen("sqlite3", "testdata/jepptest.db")
	log.Debugf("NEW -- %v", db != nil)
	log.Debug(os.Getenv("PWD"))

	return &JeppDB{DB: db}
}
