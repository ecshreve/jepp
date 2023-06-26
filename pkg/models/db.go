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

	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbaddr := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	cfg := mysql.NewConfig()
	cfg.User = dbuser
	cfg.Passwd = dbpass
	cfg.Net = "tcp"
	cfg.Addr = dbaddr
	cfg.DBName = dbname
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true
	cfg.MaxAllowedPacket = 64 << 20

	// Get a database handle.
	db = sqlx.MustOpen("mysql", cfg.FormatDSN())
	log.Debug("created new database handle", "db", db)
	return db
}
