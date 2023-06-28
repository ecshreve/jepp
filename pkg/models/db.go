package models

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

	cfg := mysql.NewConfig()
	cfg.User = "jepp"
	cfg.Passwd = "jepp"
	cfg.DBName = "jeppdb"
	cfg.Net = "tcp"
	cfg.Addr = "db:3306"
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true
	cfg.MaxAllowedPacket = 64 << 20

	// Get a database handle.
	db = sqlx.MustOpen("mysql", cfg.FormatDSN())
	log.Debug("created new database handle", "db", db)
	return db
}
