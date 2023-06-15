package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

const TIME_FORMAT = "Monday, January 2, 2006"

var re = regexp.MustCompile(`.*#([0-9]+) - (.*)$`)

type Game struct {
	GameID   int64     `db:"game_id" json:"gameId"`
	ShowNum  int64     `db:"show_num" json:"showNum"`
	GameDate time.Time `db:"game_date" json:"gameDate"`
}

func (g Game) String() string {
	return fmt.Sprintf("ID: %d -- %d - %s", g.GameID, g.ShowNum, g.GameDate.Format(TIME_FORMAT))
}

func (db *JeppDB) InsertGame(g *Game) error {
	if g == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO game (game_id, show_num, game_date) VALUES (:game_id, :show_num, :game_date)", g)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback game insert: %v", g)
		}
	}

	if err := tx.Commit(); err == nil {
		slog.Info("inserted game", "game", g)
	}
	return nil
}

// GetAllGames returns all games in the database.
func (db *JeppDB) GetAllGames() ([]*Game, error) {
	var games []*Game
	if err := db.Select(&games, "SELECT * FROM game"); err != nil {
		return nil, oops.Wrapf(err, "could not get all games")
	}

	if len(games) == 0 {
		return nil, nil
	}

	return games, nil
}
