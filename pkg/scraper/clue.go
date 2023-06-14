package scraper

import (
	"fmt"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

type Clue struct {
	ClueID   string `db:"clue_id"`
	GameID   int64  `db:"game_id"`
	Category string `db:"category"`
	Question string `db:"question"`
	Answer   string `db:"answer"`
}

// String implements fmt.Stringer.
func (c *Clue) String() string {
	return fmt.Sprintf("%d - %s", c.GameID, c.ClueID)
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
