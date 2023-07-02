package server

//go:generate go run github.com/99designs/gqlgen generate
import (
	"log"

	"github.com/ecshreve/jepp/graphql/models"
	gqlserver "github.com/ecshreve/jepp/graphql/server/generated"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func New() *gqlserver.Config {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Season{}, &models.Category{}, &models.Game{}, &models.Clue{})

	c := &gqlserver.Config{
		Resolvers: &Resolver{
			DB: db,
		},
	}

	return c
}
