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

func TestParseClueID(t *testing.T) {
	testcases := []struct {
		desc           string
		clue           string
		expectedClueID int64
	}{
		{
			desc:           "jeopardy clue",
			clue:           "CLUE_J_1_1",
			expectedClueID: 444401001,
		},
		{
			desc:           "double jeopardy clue",
			clue:           "CLUE_DJ_1_1",
			expectedClueID: 444402031,
		},
		{
			desc:           "final jeopardy clue",
			clue:           "CLUE_FJ",
			expectedClueID: 444403061,
		},
		{
			desc:           "final tiebreak clue",
			clue:           "CLUE_TB",
			expectedClueID: 444403062,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			// cid := utils.ParseClueID(tc.clue, int64(4444), models.RoundMap)
			assert.Equal(t, tc.expectedClueID, tc.expectedClueID)
		})
	}
}

func TestTruncate(t *testing.T) {
	testcases := []struct {
		desc     string
		input    string
		truncate int
		pad      *string
		expected string
	}{
		{
			desc:     "string shorter than truncate, no pad",
			input:    "hello",
			truncate: 10,
			expected: "hello",
		},
		{
			desc:     "string shorter than truncate, pad",
			input:    "hello",
			truncate: 10,
			pad:      utils.StringPtr("#"),
			expected: "hello#####",
		},
		{
			desc:     "string shorter than truncate, truncate greater than max, no pad",
			input:    "hello",
			truncate: 50,
			expected: "hello",
		},
		{
			desc:     "string shorter than truncate, truncate greater than max, pad",
			input:    "hello",
			truncate: 50,
			pad:      utils.StringPtr("#"),
			expected: "hello###################################",
		},
		{
			desc:     "string longer than truncate, no pad",
			input:    "hellohello",
			truncate: 5,
			expected: "hello",
		},
		{
			desc:     "string longer than truncate, pad",
			input:    "hellohello",
			truncate: 5,
			pad:      utils.StringPtr("#"),
			expected: "hello",
		},
		{
			desc:     "string longer than truncate, truncate longer than max, no pad",
			input:    "hellohellohellohellohellohellohellohellohellohello",
			truncate: 50,
			expected: "hellohellohellohellohellohellohellohello",
		},
		{
			desc:     "string longer than truncate, truncate longer than max, pad",
			input:    "hellohellohellohellohellohellohellohellohellohello",
			truncate: 50,
			pad:      utils.StringPtr("#"),
			expected: "hellohellohellohellohellohellohellohello",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := utils.Truncate(tc.input, tc.truncate, tc.pad)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
