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
