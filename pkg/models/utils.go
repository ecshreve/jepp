package models

import (
	"regexp"
	"strconv"
	"strings"
)

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
	clean := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(s, "")
	clean += "0000000000000000"
	clean = strings.ToUpper(clean)

	return clean[:16]
}
