package models

import (
	"fmt"

	"github.com/ecshreve/jepp/pkg/utils"
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

var RoundMap = map[string]int{
	"J":  int(Jeopardy),
	"DJ": int(DoubleJeopardy),
	"FJ": int(FinalJeopardy),
	"TB": int(FinalJeopardy),
}

// Clue represents a jeopardy clue in the database.
type Clue struct {
	ClueID     int64  `db:"clue_id" json:"clueId" example:"804002032"`
	GameID     int64  `db:"game_id" json:"gameId" example:"8040"`
	CategoryID int64  `db:"category_id" json:"categoryId" example:"3462"`
	Question   string `db:"question" json:"question" example:"This is the question."`
	Answer     string `db:"answer" json:"answer" example:"This is the answer."`
}

func (c *Clue) Dump() []string {
	ret := make([]string, 5)
	ret[0] = fmt.Sprintf("%d", c.ClueID)
	ret[1] = fmt.Sprintf("%d", c.GameID)
	ret[2] = fmt.Sprintf("%d", c.CategoryID)
	ret[3] = c.Question
	ret[4] = c.Answer

	return ret
}

// String implements the Stringer interface for the Clue type.
func (c *Clue) String() string {
	return fmt.Sprintf("%d - %d -- Q: %s - A: %s", c.ClueID, c.GameID, utils.Truncate(c.Question, 20, nil), utils.Truncate(c.Answer, 20, nil))
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
	Page       int64
	Limit      int64
}

// ListClues returns a list of clues in the database, defaults to returning
// values ordered by game id descending.
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
	query += fmt.Sprintf(" ORDER BY clue_id DESC LIMIT %d", params.Limit)

	var cc []Clue
	if err := db.Select(&cc, query); err != nil {
		return nil, oops.Wrapf(err, "could not list clues")
	}

	return cc, nil
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
	if err := db.Get(&c, "SELECT * FROM clue ORDER BY RANDOM() LIMIT 1"); err != nil {
		return nil, oops.Wrapf(err, "getting random clue")
	}

	return &c, nil
}

// GetRandomClueMany returns `limit` random clues from the database.
func GetRandomClueMany(limit int64) ([]Clue, error) {
	query := fmt.Sprintf("SELECT * FROM clue ORDER BY RANDOM() LIMIT %d", limit)

	var cc []Clue
	if err := db.Select(&cc, query); err != nil {
		return nil, oops.Wrapf(err, "getting random clues")
	}

	return cc, nil
}

// CountClues returns the number of clues in the database.
func CountClues() (int64, error) {
	var count int64
	if err := db.Get(&count, "SELECT COUNT(*) FROM clue"); err != nil {
		return -1, oops.Wrapf(err, "could not get count of clues")
	}

	return count, nil
}
