package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

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
func (db *JeppDB) InsertGame(g *Game) error {
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

// ListGames returns a list of games in the database, defaults to returning
// values ordered by game date, with most recent first.
func (db *JeppDB) ListGames(params *PaginationParams) ([]*Game, error) {
	pageSize := params.PageSize
	offset := (params.Page - 1) * params.PageSize

	var games []*Game
	if err := db.Select(&games, "SELECT * FROM game ORDER BY game_date DESC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not list games")
	}

	if len(games) == 0 {
		return nil, nil
	}

	return games, nil
}

// GetGamesBySeason returns a list of games in the database for a given season.
func (db *JeppDB) GetGamesBySeason(seasonID int64) ([]*Game, error) {
	var games []*Game
	if err := db.Select(&games, "SELECT * FROM game WHERE season_id = ?", seasonID); err != nil {
		return nil, oops.Wrapf(err, "could not get games for season %d", seasonID)
	}

	if len(games) == 0 {
		return nil, nil
	}

	return games, nil
}

// GetGame returns a single game from the database.
func (db *JeppDB) GetGame(gameID int64) (*Game, error) {
	var game Game
	if err := db.Get(&game, "SELECT * FROM game WHERE game_id = ?", gameID); err != nil {
		return nil, oops.Wrapf(err, "could not get game %d", gameID)
	}

	return &game, nil
}

// GetGame returns a single game from the database.
func (db *JeppDB) GetRandomGame() (*Game, error) {
	var allGameIDs []int64
	if err := db.Select(&allGameIDs, "SELECT game_id FROM game"); err != nil {
		return nil, oops.Wrapf(err, "getting gid")
	}

	gameID := allGameIDs[rand.Int63n(int64(len(allGameIDs)))]
	return db.GetGame(gameID)
}
