package models

import (
	"fmt"
	"time"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Season struct {
	SeasonID  int64     `db:"season_id" json:"seasonId" example:"38"`
	StartDate time.Time `db:"start_date" json:"startDate" example:"2001-09-03T00:00:00Z"`
	EndDate   time.Time `db:"end_date" json:"endDate" example:"2002-07-26T00:00:00Z"`
}

// String implements fmt.Stringer for the Season type.
func (s Season) String() string {
	return fmt.Sprintf("Season %d :: ", s.SeasonID) + s.StartDate.Format(TIME_FORMAT) + " - " + s.EndDate.Format(TIME_FORMAT)
}

// InsertSeason inserts a season into the database.
func (db *JeppDB) InsertSeason(s *Season) error {
	if s == nil {
		return nil
	}

	tx := db.MustBegin()
	_, err := db.NamedExec("INSERT INTO season (season_id, start_date, end_date) VALUES (:season_id, :start_date, :end_date)", s)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return oops.Wrapf(rollbackErr, "could not rollback season insert: %v", s)
		}
	}

	if err := tx.Commit(); err == nil {
		log.Info("inserted season", "season", s)
	}
	return nil
}

// GetSeasons returns a list of seasons in the database, defaults to returning
// values ordered by season id, with most recent first.
func (db *JeppDB) GetSeasons(params *PaginationParams) ([]*Season, error) {
	pageSize := params.PageSize
	offset := (params.Page - 1) * params.PageSize

	var seasons []*Season
	if err := db.Select(&seasons, "SELECT * FROM season ORDER BY season_id DESC LIMIT ? OFFSET ?", pageSize, offset); err != nil {
		return nil, oops.Wrapf(err, "could not list seasons")
	}

	if len(seasons) == 0 {
		return nil, nil
	}
	return seasons, nil
}
