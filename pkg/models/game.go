package models

import (
	"fmt"
	"time"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// TIME_FORMAT is the format used for friendly date strings. This is the format
// used by j-archive.
const TIME_FORMAT = "Monday, January 2, 2006"

// Game represents a single game of Jeopardy.
type Game struct {
	GameID    int64     `db:"game_id" json:"gameId" example:"8040"`
	SeasonID  int64     `db:"season_id" json:"seasonId" example:"38"`
	ShowNum   int64     `db:"show_num" json:"showNum" example:"4532"`
	GameDate  time.Time `db:"game_date" json:"gameDate" example:"2019-01-01"`
	TapedDate time.Time `db:"taped_date" json:"tapedDate" example:"2019-01-01"`
}

// String implements fmt.Stringer for the Game type.
func (g Game) String() string {
	return fmt.Sprintf("ID: %d -- Show: %d - Aired: %s", g.GameID, g.ShowNum, g.GameDate.Format(TIME_FORMAT))
}

// InsertGame inserts a game into the database.
func InsertGame(g *Game) error {
	if g == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO game (game_id, season_id, show_num, game_date, taped_date) VALUES (:game_id, :season_id, :show_num, :game_date, :taped_date)", g)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback game insert: %v", g)
		}
	}

	if err := tx.Commit(); err == nil {
		log.Infof("inserted game %+v", g)
	}

	return nil
}

// GetGames returns a list of games in the database, defaults to returning
// values ordered by game date, with most recent first, limited to 100 results.
//
// TODO: have this take a "lastClueID" arg or something for dumb pagination.
func GetGames() ([]Game, error) {
	games := []Game{}
	if err := db.Select(&games, "SELECT * FROM game ORDER BY game_date DESC LIMIT 100"); err != nil {
		return nil, oops.Wrapf(err, "could not list games")
	}

	return games, nil
}

// GetGamesBySeason returns a list of games in the database for a given season.
func GetGamesBySeason(seasonID int64) ([]Game, error) {
	query := fmt.Sprintf("SELECT * FROM game WHERE season_id=%d ORDER BY game_date DESC", seasonID)

	games := []Game{}
	if err := db.Select(&games, query); err != nil {
		return nil, oops.Wrapf(err, "could not get games for season %d", seasonID)
	}

	return games, nil
}

// GetGame returns a single game from the database.
func GetGame(gameID int64) (*Game, error) {
	query := fmt.Sprintf("SELECT * FROM game WHERE game_id=%d LIMIT 1", gameID)

	g := Game{}
	if err := db.Get(&g, query); err != nil {
		return nil, oops.Wrapf(err, "could not get game %d", gameID)
	}

	return &g, nil
}

// GetRandomGame returns a single game from the database.
func GetRandomGame() (*Game, error) {
	g := Game{}
	if err := db.Get(&g, "SELECT * FROM game ORDER BY RAND() LIMIT 1"); err != nil {
		return nil, oops.Wrapf(err, "getting gid")
	}

	return &g, nil
}
