package models

import (
	"fmt"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

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

type Clue struct {
	ClueID   int64  `db:"clue_id" json:"clueId"`
	GameID   int64  `db:"game_id" json:"gameId"`
	Category string `db:"category" json:"category"`
	Question string `db:"question" json:"question"`
	Answer   string `db:"answer" json:"answer"`
}

// String implements fmt.Stringer.
func (c *Clue) String() string {
	return fmt.Sprintf("%d - %d - %s", c.GameID, c.ClueID, c.Category)
}

func (db *JeppDB) InsertClue(c *Clue) error {
	if c == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO clue (clue_id, game_id, category, question, answer) VALUES (:clue_id, :game_id, :category, :question, :answer)", c)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback clue insert: %v", c)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Debug("inserted clue", "clue", c)
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
	if err := db.Select(&clues, "SELECT * FROM clue WHERE game_id = ? ORDER BY clueID ASC", gameId); err != nil {
		return nil, oops.Wrapf(err, "could not get clues for game_id %s", gameId)
	}

	if len(clues) == 0 {
		return nil, nil
	}

	return clues, nil
}
