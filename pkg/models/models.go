package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type JeppDB struct {
	*sqlx.DB
}

func NewJeppDB(dbfile string) *JeppDB {
	db := sqlx.MustOpen("sqlite3", dbfile)
	return &JeppDB{DB: db}
}
