package models

import (
	"regexp"
	"strings"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

type Category struct {
	CategoryID string `db:"category_id" json:"categoryId"`
	Name       string `db:"name" json:"name"`
	GameID     int64  `db:"game_id" json:"gameId"`
}

type CategoryCount struct {
	CategoryID   string `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	GameCount    int    `db:"game_count" json:"gameCount"`
	ClueCount    int    `db:"clue_count" json:"clueCount"`
}

func GetCategoryID(s string) string {
	clean := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(s, "")
	clean += "0000000000000000"
	clean = strings.ToUpper(clean)
	return clean[:16]
}

// InsertCategory inserts a category into the database.
func (db *JeppDB) InsertCategory(c *Category) error {
	if c == nil {
		return nil
	}

	c.CategoryID = GetCategoryID(c.Name)

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO category (category_id, name, game_id) VALUES (:category_id, :name, :game_id)", c)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback category insert: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Info("inserted category", "category", c)
	}
	return nil
}

// UpdateCategory updates a category in the database.
func (db *JeppDB) UpdateCategory(c *Category) error {
	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	if _, err := tx.NamedExec("UPDATE category SET category_id=:category_id WHERE name=:name AND game_id=:game_id", c); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback category update: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Info("updated category", "category", c)
	}

	return nil
}

// ListCategories returns all categories in the database.
func (db *JeppDB) ListCategories(params *PaginationParams) ([]*CategoryCount, error) {
	if params == nil {
		params = &PaginationParams{Page: 1, PageSize: 10}
	}

	pageSize := params.PageSize
	offset := params.Page * params.PageSize

	var categories []*CategoryCount
	if err := db.Select(&categories, "SELECT * FROM category_counts ORDER BY category_id ASC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not get all categories")
	}

	if len(categories) == 0 {
		return nil, nil
	}

	return categories, nil
}

// GetCategoriesForGame returns all categories for a given game.
func (db *JeppDB) GetCategoriesForGame(gameID int64) ([]*Category, error) {
	var categories []*Category
	if err := db.Select(&categories, "SELECT * FROM category WHERE game_id=?", gameID); err != nil {
		return nil, oops.Wrapf(err, "could not get categories for game %d", gameID)
	}

	if len(categories) == 0 {
		return nil, nil
	}

	return categories, nil
}
