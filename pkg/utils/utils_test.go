package utils_test

import (
	"testing"

	"github.com/ecshreve/jepp/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseRoundAndColumn(t *testing.T) {
	testcases := []struct {
		desc           string
		clue           string
		expectedRound  int64
		expectedColumn int64
	}{
		{
			desc:           "jeopardy clue",
			clue:           "CLUE_J_1_1",
			expectedRound:  1,
			expectedColumn: 1,
		},
		{
			desc:           "double jeopardy clue",
			clue:           "CLUE_DJ_1_1",
			expectedRound:  2,
			expectedColumn: 1,
		},
		{
			desc:           "final jeopardy clue",
			clue:           "CLUE_FJ",
			expectedRound:  3,
			expectedColumn: 1,
		},
		{
			desc:           "final tiebreak clue",
			clue:           "CLUE_TB",
			expectedRound:  3,
			expectedColumn: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			rd, col := utils.ParseRoundAndColumn(tc.clue)
			assert.Equal(t, tc.expectedRound, rd)
			assert.Equal(t, tc.expectedColumn, col)
		})
	}
}

func TestGetCategoryIDInt(t *testing.T) {
	testcases := []struct {
		desc     string
		clue     string
		gameId   int64
		expected int64
	}{
		{
			desc:     "jeopardy clue",
			clue:     "CLUE_J_1_1",
			gameId:   1111,
			expected: 111191001,
		},
		{
			desc:     "double jeopardy clue",
			clue:     "CLUE_DJ_1_1",
			gameId:   1111,
			expected: 111192001,
		},
		{
			desc:     "final jeopardy clue",
			clue:     "CLUE_FJ",
			gameId:   1111,
			expected: 111193001,
		},
		{
			desc:     "final tiebreak clue",
			clue:     "CLUE_TB",
			gameId:   1111,
			expected: 111193002,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			rd, col := utils.ParseRoundAndColumn(tc.clue)
			actual := utils.GetCategoryIDInt(tc.gameId, rd, col)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}

}
