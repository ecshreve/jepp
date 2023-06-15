package models

import (
	"strconv"
	"strings"
)

type PaginationParams struct {
	Page     int
	PageSize int
}

// GetClueID converts a clue string to a clue ID.
//
// Clue strings have the format "clue_DJ_1_2", "clue_FJ"
func GetClueID(clueString string, gameId int64) int64 {
	baseVal := gameId * 100000
	tokens := strings.Split(clueString, "_")
	if len(tokens) == 2 {
		if tokens[1] == "FJ" {
			return baseVal + 3061
		}
		return baseVal + 3062
	}

	round := RoundMap[tokens[1]]
	baseVal = baseVal + (int64(round) * 1000)

	column, _ := strconv.ParseInt(tokens[2], 10, 64)
	row, _ := strconv.ParseInt(tokens[3], 10, 64)

	return baseVal + ((int64(round) - 1) * 30) + (((column - 1) * 5) + row)
}
