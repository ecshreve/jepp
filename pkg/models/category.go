package models

import (
	"fmt"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// Category represents a jeopardy category in the database.
type Category struct {
	CategoryID int64  `db:"category_id" json:"categoryId" example:"765"`
	Name       string `db:"name" json:"name" example:"State Capitals"`
}

// GetCategoryGameCount returns the number of games a category has appeared in.
func GetCategoryGameCount(categoryID int64) (int64, error) {

	var count int64

	err := db.Get(&count, "SELECT COUNT(DISTINCT game_id) FROM clue WHERE category_id = ?", categoryID)
	if err != nil {
		return 0, oops.Wrapf(err, "could not get category game count")
	}

	return count, nil
}

// GetCategoryClueCount returns the number of clues a category has appeared in.
func GetCategoryClueCount(categoryID int64) (int64, error) {

	var count int64

	err := db.Get(&count, "SELECT COUNT(*) FROM clue WHERE category_id = ?", categoryID)
	if err != nil {
		return 0, oops.Wrapf(err, "could not get category clue count")
	}

	return count, nil
}

// InsertCategory inserts a category into the database.
func InsertCategory(name string) (*Category, error) {

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
		log.Debugf("inserted category -- %+v", cat)
	} else {
		return nil, oops.Wrapf(err, "could not insert category: %v", name)
	}
	return cat, nil
}

// UpdateCategory updates a category in the database.
func UpdateCategory(c *Category) error {

	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	if _, err := tx.NamedExec("UPDATE category SET name=:name WHERE category_id=:category_id", c); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback category update: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		log.Infof("updated category %+v", c)
	}

	return nil
}

// GetCategories returns all categories in the database.
func GetCategories() ([]Category, error) {
	var categories []Category
	if err := db.Select(&categories, "SELECT * FROM category ORDER BY category_id ASC LIMIT 100"); err != nil {
		return nil, oops.Wrapf(err, "could not get all categories")
	}

	return categories, nil
}

// GetCategory returns the category with the given ID.
func GetCategory(categoryID int64) (*Category, error) {
	query := fmt.Sprintf("SELECT * FROM category WHERE category_id=%d ORDER BY category_id DESC LIMIT 1", categoryID)

	c := Category{}
	if err := db.Get(&c, query); err != nil {
		return nil, oops.Wrapf(err, "could not get category for id %d", categoryID)
	}

	return &c, nil
}

func GetCategoryByName(categoryName string) (*Category, error) {
	query := fmt.Sprintf("SELECT * FROM category WHERE name='%s' ORDER BY category_id DESC LIMIT 1", categoryName)

	c := Category{}
	if err := db.Get(&c, query); err != nil {
		return nil, oops.Wrapf(err, "could not get category for name %s", categoryName)
	}

	log.Debugf("category: %+v", c)
	return &c, nil
}

// GetRandomCategory returns a single category from the database.
func GetRandomCategory() (*Category, error) {
	c := Category{}
	if err := db.Get(&c, "SELECT * FROM category ORDER BY RAND() LIMIT 1"); err != nil {
		return nil, oops.Wrapf(err, "getting random category")
	}

	return &c, nil
}

// GetRandomCategoryMany returns `count` random categories from the database.
func GetRandomCategoryMany(count int64) ([]Category, error) {
	query := fmt.Sprintf("SELECT * FROM category ORDER BY RAND() LIMIT %d", count)

	var cc []Category
	if err := db.Select(&cc, query); err != nil {
		return nil, oops.Wrapf(err, "getting random categories")
	}

	return cc, nil
}

// GetCategoriesForGame returns all categories for a given game.
func GetCategoriesForGame(gameID int64) ([]Category, error) {
	var categories []Category
	if err := db.Select(&categories, "SELECT clue.category_id, category.name FROM clue JOIN category ON clue.category_id = category.category_id WHERE game_id=? GROUP BY category_id", gameID); err != nil {
		return nil, oops.Wrapf(err, "could not get categories for game %d", gameID)
	}

	return categories, nil
}
