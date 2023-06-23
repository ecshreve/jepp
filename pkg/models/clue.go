package models

import (
	"fmt"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
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
	CategoryID int64  `db:"category_id" json:"categoryId" example:"804092001"`
	Question   string `db:"question" json:"question" example:"This is the question."`
	Answer     string `db:"answer" json:"answer" example:"This is the answer."`
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

// String implements fmt.Stringer.
func (c *Clue) String() string {
	return fmt.Sprintf("%d - %d -- Q: %s - A: %s", c.ClueID, c.GameID, truncate(c.Question, 20), truncate(c.Answer, 20))
}

// InsertClue inserts a clue into the database.
func InsertClue(c *Clue) error {

	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO clue (clue_id, game_id, category_id, question, answer) VALUES (:clue_id, :game_id, :category_id, :question, :answer)", c)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback clue insert: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		log.Debugf("inserted clue %+v", c)
	}
	return nil
}

// UpdateClue updates a category in the database.
func UpdateClue(c *Clue) error {
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
		log.Debugf("updated clue %+v", c)
	}

	return nil
}

type CluesParams struct {
	GameID     int64
	CategoryID int64
}

// ListClues returns a list of clues in the database, defaults to returning
// values ordered by game date, with most recent first.
func GetClues(params CluesParams) ([]Clue, error) {
	query := "SELECT * FROM clue"

	if params.GameID != 0 {
		query += fmt.Sprintf(" WHERE game_id=%d", params.GameID)
	}

	if params.CategoryID != 0 {
		if params.GameID != 0 {
			query += " AND"
		} else {
			query += " WHERE"
		}

		query += fmt.Sprintf(" category_id=%d", params.CategoryID)
	}
	query += " ORDER BY clue_id DESC LIMIT 100"
	log.Debug("query: ", query)

	var clues []Clue
	if err := db.Select(&clues, query); err != nil {
		return nil, oops.Wrapf(err, "could not list clues")
	}

	return clues, nil
}

// GetClue returns a single clue from the database.
func GetClue(clueID int64) (*Clue, error) {
	query := fmt.Sprintf("SELECT * FROM clue WHERE clue_id=%d ORDER BY clue_id DESC LIMIT 1", clueID)

	c := Clue{}
	if err := db.Get(&c, query); err != nil {
		return nil, oops.Wrapf(err, "could not get clue %d", clueID)
	}

	return &c, nil
}

// GetClue returns a single clue from the database.
func GetRandomClue() (*Clue, error) {
	c := Clue{}
	if err := db.Get(&c, "SELECT * FROM clue ORDER BY RAND() LIMIT 1"); err != nil {
		return nil, oops.Wrapf(err, "getting random clue")
	}

	return &c, nil
}

// NumClues returns the number of clues in the database.
func NumClues() (int64, error) {
	var count int64
	if err := db.Get(&count, "SELECT COUNT(*) FROM clue"); err != nil {
		return 0, oops.Wrapf(err, "could not get count of clues")
	}

	return count, nil
}
