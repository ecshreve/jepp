package models

import (
	"fmt"
	"time"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

const TIME_FORMAT = "Monday, January 2, 2006"

// Game represents a single game of Jeopardy.
type Game struct {
	GameID   int64     `db:"game_id" json:"gameId" example:"8040"`
	ShowNum  int64     `db:"show_num" json:"showNum" example:"4532"`
	GameDate time.Time `db:"game_date" json:"gameDate" example:"2019-01-01"`
}

// String implements fmt.Stringer for the Game type.
func (g Game) String() string {
	return fmt.Sprintf("ID: %d -- %d - %s", g.GameID, g.ShowNum, g.GameDate.Format(TIME_FORMAT))
}

// InsertGame inserts a game into the database.
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

// ListGames returns a list of games in the database, defaults to returning
// values ordered by game date, with most recent first.
func (db *JeppDB) ListGames(params *PaginationParams) ([]*Game, error) {
	if params == nil {
		params = &PaginationParams{Page: 1, PageSize: 10}
	}

	pageSize := params.PageSize
	offset := params.Page * params.PageSize

	var games []*Game
	if err := db.Select(&games, "SELECT * FROM game ORDER BY game_date DESC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not list games")
	}

	if len(games) == 0 {
		return nil, nil
	}

	return games, nil
}
