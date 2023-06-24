package utils

import (
	"strconv"
	"strings"
)

// ParseRoundAndColumn parses a clue string and returns the round and column.
func ParseRoundAndColumn(clueString string) (int64, int64) {
	tokens := strings.Split(clueString, "_")

	var rd int64
	if tokens[1] == "J" {
		rd = 1
	} else if tokens[1] == "DJ" {
		rd = 2
	} else if tokens[1] == "FJ" {
		return 3, 1
	} else if tokens[1] == "TB" {
		return 3, 2
	}

	if len(tokens) != 4 {
		return -1, -1
	}

	col, _ := strconv.ParseInt(tokens[2], 10, 64)
	return rd, col
}

// ParseClueID converts a clue string to a clue ID.
// Clue strings have the format "clue_DJ_1_2", "clue_FJ" and the parsed int64
// is of the form <game_id>0<round>0<clue_index>.
func ParseClueID(clueString string, gameId int64, rm map[string]int) int64 {
	baseVal := gameId * 100000
	tokens := strings.Split(clueString, "_")
	if len(tokens) == 2 {
		if tokens[1] == "FJ" {
			return baseVal + 3061
		}
		return baseVal + 3062
	}

	// TODO: this is hacky, and I forget how it works
	round := rm[tokens[1]]                           // val = 804000000 round = DJ = 2
	baseVal = baseVal + (int64(round) * 1000)        // val = 804002000
	column, _ := strconv.ParseInt(tokens[2], 10, 64) // column = 1, row = 2
	row, _ := strconv.ParseInt(tokens[3], 10, 64)    // val = 804002032
	return baseVal + ((int64(round) - 1) * 30) + (((column - 1) * 5) + row)
}
