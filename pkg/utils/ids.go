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

// GetCategoryIDInt gets a category ID from a clue string and game ID.
func GetCategoryIDInt(gameId, round, column int64) int64 {
	baseVal := gameId*10 + 9
	baseVal = baseVal * 10000
	baseVal = baseVal + (round * 1000) + column
	return baseVal
	// tokens := strings.Split(clueString, "_")
	// if len(tokens) == 2 {
	// 	if tokens[1] == "FJ" {
	// 		return baseVal + 3001
	// 	}
	// 	return baseVal + 3002
	// }

	// if len(tokens) != 4 {
	// 	return -1
	// }

	// if tokens[1] == "J" {
	// 	baseVal += 1000
	// } else if tokens[1] == "DJ" {
	// 	baseVal += 2000
	// }

	// column, _ := strconv.ParseInt(tokens[2], 10, 64)
	// return baseVal + column
}
