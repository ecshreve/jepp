package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var db *sqlx.DB

// JeppDB is a wrapper around sqlx.DB.
type JeppDB struct {
	Conn *sqlx.DB
}

// GetDBHandler returns a new database handle.
func GetDBHandle() *sqlx.DB {
	if db != nil {
		log.Debug("using existing database handle", "db", db)
		return db
	}

	// Get a database handle.
	db = sqlx.MustConnect("sqlite3", "../../data/jepp.db")
	log.Debug("created new database handle", "db", db)
	return db
}
