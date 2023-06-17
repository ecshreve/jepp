package models

import (
	"math/rand"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// Category represents a jeopardy category in the database.
type Category struct {
	CategoryID int64  `db:"category_id" json:"categoryId" example:"765"`
	Name       string `db:"name" json:"name" example:"Category Name"`
}

// CategoryCount represents a the number of times a category has appeared in
// different games and clues.
type CategoryCount struct {
	CategoryID   int64  `db:"category_id" json:"categoryId" example:"804001001"`
	CategoryName string `db:"category_name" json:"categoryName" example:"Category Name"`
	GameCount    int    `db:"game_count" json:"gameCount" example:"1"`
	ClueCount    int    `db:"clue_count" json:"clueCount" example:"5"`
}

// InsertCategory inserts a category into the database.
func (db *JeppDB) InsertCategory(name string) (*Category, error) {
	if name == "" {
		return nil, oops.Errorf("cannot insert empty category")
	}

	tx := db.MustBegin()
	res, err := db.Exec("INSERT INTO category (name) VALUES (?)", name)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, oops.Wrapf(rollbackErr, "could not rollback category insert: %v", name)
		}
	}

	var cat *Category
	if err := tx.Commit(); err == nil {
		lid, _ := res.LastInsertId()
		cat = &Category{CategoryID: lid, Name: name}
		log.Infof("inserted category -- %+v", cat)
	} else {
		return nil, oops.Wrapf(err, "could not insert category: %v", name)
	}
	return cat, nil
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
		log.Info("updated category", "category", c)
	}

	return nil
}

// ListCategories returns all categories in the database.
func (db *JeppDB) ListCategories(params *PaginationParams) ([]*Category, error) {
	if params == nil {
		params = &PaginationParams{Page: 0, PageSize: 10}
	}

	pageSize := params.PageSize
	offset := params.Page * params.PageSize

	var categories []*Category
	if err := db.Select(&categories, "SELECT * FROM category ORDER BY category_id ASC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not get all categories")
	}

	if len(categories) == 0 {
		return nil, nil
	}

	return categories, nil
}

// func (db *JeppDB) GetCategoryCount(categoryID int64) (*CategoryCount, error) {
// 	var category CategoryCount
// 	if err := db.Get(&category, "SELECT * FROM category_counts WHERE category_id=?", categoryID); err != nil {
// 		return nil, oops.Wrapf(err, "could not get category for id %d", categoryID)
// 	}
// 	return &category, nil
// }

// GetCategory returns the category with the given ID.
func (db *JeppDB) GetCategory(categoryID int64) (*Category, error) {
	var category Category
	if err := db.Get(&category, "SELECT * FROM category WHERE category_id=?", categoryID); err != nil {
		return nil, oops.Wrapf(err, "could not get category for id %d", categoryID)
	}

	return &category, nil
}

func (db *JeppDB) GetCategoryByName(categoryName string) (*Category, error) {
	var category Category
	if err := db.Get(&category, "SELECT * FROM category WHERE name=?", categoryName); err != nil {
		return nil, oops.Wrapf(err, "could not get category for name %s", categoryName)
	}

	return &category, nil
}

// GetRandomCategory returns a single category from the database.
func (db *JeppDB) GetRandomCategory() (*Category, error) {
	var allCategoryIDs []int64
	if err := db.Select(&allCategoryIDs, "SELECT category_id FROM category"); err != nil {
		return nil, oops.Wrapf(err, "getting category ids")
	}

	categoryID := allCategoryIDs[rand.Int63n(int64(len(allCategoryIDs)))]
	return db.GetCategory(categoryID)
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
