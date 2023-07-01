package graph

//go:generate go run generate.go
import (
	"log"

	"github.com/ecshreve/jepp/graph/common"
	graph "github.com/ecshreve/jepp/graph/generated"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func New() *graph.Config {
	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	c := &graph.Config{
		Resolvers: &Resolver{
			DB: db,
		},
	}
	return c
}
