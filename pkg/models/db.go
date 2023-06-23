package models

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// JeppDB is a wrapper around sqlx.DB.
type JeppDB struct {
	Conn *sqlx.DB
}

// NewDB returns a new database handle.
func GetDBHandle() *sqlx.DB {
	if db != nil {
		return db
	}

	dbname := "jeppdb"
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

	// Capture connection properties.
	// cfg = mysql.Config{
	// 	User:                 dbuser,
	// 	Passwd:               dbpass,
	// 	Net:                  "tcp",
	// 	Addr:                 dbaddr,
	// 	DBName:               dbname,
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// 	MaxAllowedPacket: 64 << 20,
	// }

	// Get a database handle.
	dbx := sqlx.MustOpen("mysql", cfg.FormatDSN())
	db = dbx
	return db
}

// InitDB initializes the database.
//
// Not currently used. I manually create tables in Adminer.
// func (db *JeppDB) initDB() {
// 	db.MustExec(GAME_SCHEMA)
// 	db.MustExec(CLUE_SCHEMA)
// 	db.MustExec(CATEGORY_SCHEMA)
// }
