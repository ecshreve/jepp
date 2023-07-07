package utils

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ecshreve/jepp/internal/ent"
	log "github.com/sirupsen/logrus"
)

func InitDB() (*ent.Client, *sql.DB) {
	log.Info("Initializing database...")
	// Create ent sql driver.
	dd, err := entsql.Open(dialect.SQLite, "data/sqlite/jepp.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	// Create ent client.
	cc := ent.NewClient(ent.Driver(dd))
	if err := cc.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	return cc, dd.DB()
}
