package models

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Option struct {
	OptionKey string
	OptionVal string
	Selected  bool
}

type NavLinks struct {
	PrevClue string
	NextClue string
}

type Options struct {
	ClueID          int64
	Links           NavLinks
	CategoryOptions []Option
}

// PaginationParams is a struct that holds pagination parameters.
// These apply to both API responses and database queries.
type PaginationParams struct {
	Page     int
	PageSize int
}

// ParseClueID converts a clue string to a clue ID.
//
// Clue strings have the format "clue_DJ_1_2", "clue_FJ"
func ParseClueID(clueString string, gameId int64) int64 {
	baseVal := gameId * 100000
	tokens := strings.Split(clueString, "_")
	if len(tokens) == 2 {
		if tokens[1] == "FJ" {
			return baseVal + 3061
		}
		return baseVal + 3062
	}

	// TODO: this is hacky, and I forget how it works
	round := RoundMap[tokens[1]]                     // val = 804000000 round = DJ = 2
	baseVal = baseVal + (int64(round) * 1000)        // val = 804002000
	column, _ := strconv.ParseInt(tokens[2], 10, 64) // column = 1, row = 2
	row, _ := strconv.ParseInt(tokens[3], 10, 64)    // val = 804002032
	return baseVal + ((int64(round) - 1) * 30) + (((column - 1) * 5) + row)
}

// GetCategoryID converts a category string to a category ID.
func GetCategoryID(s string) string {
	clean := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(s, "0")
	clean = strings.ToUpper(clean)
	return clean
}

// GetSeasonURL returns the scrape URL for the given season.
func GetSeasonURL(seasonID int64) string {
	return fmt.Sprintf("http://www.j-archive.com/showseason.php?season=%d", seasonID)
}
