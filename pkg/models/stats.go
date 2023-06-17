package models

import "github.com/samsarahq/go/oops"

type Stats struct {
	TotalClues      int64 `db:"total_clues" json:"totalClues" example:"100"`
	TotalGames      int64 `db:"total_games" json:"totalGames" example:"10"`
	TotalCategories int64 `db:"total_cats" json:"totalCats" example:"10"`
}

// GetStats returns the total number of clues, games, and categories in the database.
func (db *JeppDB) GetStats() (*Stats, error) {
	var stats Stats
	err := db.Get(&stats, "SELECT COUNT(*) AS total_clues, (SELECT COUNT(*) FROM game) AS total_games, (SELECT COUNT(*) FROM category) AS total_cats FROM clue")
	if err != nil {
		return nil, oops.Wrapf(err, "could not get stats")
	}
	return &stats, nil
}
