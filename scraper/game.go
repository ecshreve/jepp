package scraper

import (
	"time"

	"github.com/samsarahq/go/oops"
	"golang.org/x/exp/slog"
)

type Game struct {
	GameID   int64     `db:"game_id"`
	ShowNum  int64     `db:"show_num"`
	GameDate time.Time `db:"game_date"`
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
