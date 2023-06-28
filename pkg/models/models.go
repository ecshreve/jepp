package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type JeppDB struct {
	*sqlx.DB
	LastCategoryID int64
}

func NewJeppDB(dbfile string) *JeppDB {
	db := sqlx.MustOpen("sqlite3", dbfile)
	var lid int64
	if err := db.Get(&lid, "SELECT MAX(category_id) FROM category"); err != nil {
		log.Debugf("could not get last category id: %v", err)
		return nil
	}
	db.Get(&lid, "SELECT MAX(id) FROM category")
	return &JeppDB{DB: db, LastCategoryID: lid}
}
