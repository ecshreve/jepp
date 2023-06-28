package models

import (
	"fmt"
	"os"

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
	cfg.User = "root"
	cfg.DBName = "jeppdb"
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:3306", os.Getenv("JEPP_DB_HOST"))
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true
	cfg.MaxAllowedPacket = 64 << 20

	if os.Getenv("JEPP_ENV") == "ci" {
		cfg.Passwd = "root"
	}

	// Get a database handle.
	db = sqlx.MustOpen("mysql", cfg.FormatDSN())
	log.Debug("created new database handle", "db", db)
	return db
}
