package models

import (
	"fmt"
	"math/rand"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

// TODO: this is silly, should fix it
type Round int // 0 = Jeopardy, 1 = Double Jeopardy, 2 = Final Jeopardy

const (
	Jeopardy Round = iota + 1
	DoubleJeopardy
	FinalJeopardy
)

var RoundMap = map[string]Round{
	"J":  Jeopardy,
	"DJ": DoubleJeopardy,
	"FJ": FinalJeopardy,
	"TB": FinalJeopardy,
}

// Clue represents a jeopardy clue in the database.
type Clue struct {
	ClueID     int64  `db:"clue_id" json:"clueId" example:"804002032"`
	GameID     int64  `db:"game_id" json:"gameId" example:"8040"`
	CategoryID string `db:"category_id" json:"categoryId" example:"CATEGORYNAME0000"`
	Question   string `db:"question" json:"question" example:"This is the question."`
	Answer     string `db:"answer" json:"answer" example:"This is the answer."`
}

// String implements fmt.Stringer.
func (c *Clue) String() string {
	return fmt.Sprintf("%d - %d - %s", c.GameID, c.ClueID, c.CategoryID)
}

// InsertClue inserts a clue into the database.
func (db *JeppDB) InsertClue(c *Clue) error {
	if c == nil {
		return nil
	}

	if len(c.CategoryID) != 8 {
		c.CategoryID = GetCategoryID(c.CategoryID)
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO clue (clue_id, game_id, category_id, question, answer) VALUES (:clue_id, :game_id, :category_id, :question, :answer)", c)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback clue insert: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Info("inserted clue", "clue", c)
	}
	return nil
}

// UpdateClue updates a category in the database.
func (db *JeppDB) UpdateClue(c *Clue) error {
	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	if _, err := tx.NamedExec("UPDATE clue SET category_id=:category_id, game_id=:game_id, question=:question, answer=:answer WHERE clue_id=:clue_id", c); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback clue update: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Info("updated clue", "clue", c)
	}

	return nil
}

// GetAllClues eturns all clues in the database.
func (db *JeppDB) GetAllClues() ([]*Clue, error) {
	var clues []*Clue
	if err := db.Select(&clues, "SELECT * FROM clue"); err != nil {
		return nil, oops.Wrapf(err, "could not get all clues")
	}

	if len(clues) == 0 {
		return nil, nil
	}

	return clues, nil
}

// GetCluesForGame returns all clues for a given game.
func (db *JeppDB) GetCluesForGame(gameId string) ([]*Clue, error) {
	var clues []*Clue
	if err := db.Select(&clues, "SELECT * FROM clue WHERE game_id = ? ORDER BY clue_id ASC", gameId); err != nil {
		return nil, oops.Wrapf(err, "could not get clues for game_id %s", gameId)
	}

	if len(clues) == 0 {
		return nil, nil
	}

	return clues, nil
}

// GetCluesForCategory returns all clues for a given category.
func (db *JeppDB) GetCluesForCategory(category_id string) ([]*Clue, error) {
	var clues []*Clue
	if err := db.Select(&clues, "SELECT * FROM clue WHERE category_id = ? ORDER BY clue_id ASC", category_id); err != nil {
		return nil, oops.Wrapf(err, "could not get clues for category %s", category_id)
	}

	if len(clues) == 0 {
		return nil, nil
	}

	return clues, nil
}

// ListClues returns a list of clues in the database, defaults to returning
// values ordered by game date, with most recent first.
func (db *JeppDB) ListClues(params *PaginationParams) ([]*Clue, error) {
	if params == nil {
		params = &PaginationParams{Page: 1, PageSize: 10}
	}

	pageSize := params.PageSize
	offset := params.Page * params.PageSize

	var clues []*Clue
	if err := db.Select(&clues, "SELECT * FROM clue ORDER BY clue_id DESC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not list clues")
	}

	if len(clues) == 0 {
		return nil, nil
	}

	return clues, nil
}

// GetClue returns a single clue from the database.
func (db *JeppDB) GetClue(clueID string) (*Clue, error) {
	var clue Clue
	if err := db.Get(&clue, "SELECT * FROM clue WHERE clue_id = ?", clueID); err != nil {
		return nil, oops.Wrapf(err, "could not get clue %s", clueID)
	}

	return &clue, nil
}

// GetClue returns a single clue from the database.
func (db *JeppDB) GetRandomClue() (*Clue, error) {
	var allClueIDs []string
	if err := db.Select(&allClueIDs, "SELECT clue_id FROM clue"); err != nil {
		return nil, oops.Wrapf(err, "getting clue ids")
	}

	clueID := allClueIDs[rand.Int63n(int64(len(allClueIDs)))]
	return db.GetClue(clueID)
}
