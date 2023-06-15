package models

import (
	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

type Category struct {
	Name   string `db:"name" json:"name"`
	GameID int64  `db:"game_id" json:"gameId"`
}

// InsertCategory inserts a category into the database.
func (db *JeppDB) InsertCategory(c *Category) error {
	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO category (name, game_id) VALUES (:name, :game_id)", c)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback category insert: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Debug("inserted category", "category", c)
	}
	return nil
}

// GetAllCategories returns all categories in the database.
func (db *JeppDB) GetAllCategories() ([]*Category, error) {
	var categories []*Category
	if err := db.Select(&categories, "SELECT * FROM category"); err != nil {
		return nil, oops.Wrapf(err, "could not get all categories")
	}

	if len(categories) == 0 {
		return nil, nil
	}

	return categories, nil
}
