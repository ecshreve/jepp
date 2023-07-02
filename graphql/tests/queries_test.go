package tests

import (
	"fmt"
	"testing"
)

func TestSingleQueries(t *testing.T) {
	// Setup test environment
	env := SetupTestEnv(t)
	defer env.snap.Verify()
	defer env.Cleanup(t)

	testcases := []struct {
		q string
	}{
		{
			q: `query {
				season(seasonID: 1) {
					id			
					number
					startDate
					endDate
					games {
						id
					}
				}
			}`,
		},
		{
			q: `query {
				game(gameID: 1) {
					id
					season {
						id
					}
					show
					airDate
					tapeDate
					clues {
						id
					}
				}
			}`,
		},
		{
			q: `query {
				category(categoryID: 1) {
					id
					name
					clues {
						id
					}
				}
			}`,
		},
		{
			q: `query {
				clue(clueID: 1) {
					id
					question
					answer
					category {
						id
					}
					game {
						id
					}
				}
			}`,
		},
	}

	for i, tc := range testcases {
		desc := fmt.Sprintf("tc #%d", i)
		t.Run(desc, func(t *testing.T) {
			// Make request
			result := env.queryHelper(t, tc.q, nil)

			// Verify snapshot
			env.snap.Snapshot(desc, &result)
		})
	}
}

func TestBasicQueries(t *testing.T) {
	testcases := []struct {
		q string
	}{
		{
			q: `query {
				seasons {
					id			
					number
					startDate
					endDate
					games {
						id
					}
				}
			}`,
		},
		{
			q: `query {
				games {
					nodes {
						id
						season {
							id
						}
						show
						airDate
						tapeDate
					}
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
				}
			}`,
		},
		{
			q: `query {
				categories {
					nodes {
						id
						name
						clues {
							id
						}
					}
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
				}
			}`,
		},
		{
			q: `query {
				clues {
					nodes {
						id
						question
						answer
						category {
							id
						}
						game {
							id
						}
					}
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
				}
			}`,
		},
	}

	for i, tc := range testcases {
		desc := fmt.Sprintf("tc #%d", i)
		t.Run(desc, func(t *testing.T) {
			// Setup test environment
			env := SetupTestEnv(t)
			defer env.snap.Verify()
			defer env.Cleanup(t)

			// Make request
			result := env.queryHelper(t, tc.q, nil)

			// Verify snapshot
			env.snap.Snapshot(desc, &result)
		})
	}
}

func TestComplexQueries(t *testing.T) {
	testcases := []struct {
		q string
	}{
		{
			q: `query {
				categories {
					nodes {
						id
						name
						clues {
							id
							question
							answer
							game {
								id
								season {
									id
									number
									startDate
									endDate
								}
								show
								airDate
								tapeDate
							}
						}
					}
					pageInfo {
						startCursor
						endCursor
						hasNextPage
					}
				}
			}`,
		},
		{
			q: `query {
				seasons {
					id
					number
					startDate
					endDate
					games {
						id
						show
						airDate
						tapeDate
						clues {
							id
							question
							answer
							category {
								id
								name
							}
						}
					}
				}
			}`,
		},
	}

	for i, tc := range testcases {
		desc := fmt.Sprintf("tc #%d", i)

		t.Run(desc, func(t *testing.T) {
			// Setup test environment
			env := SetupTestEnv(t)
			defer env.snap.Verify()
			defer env.Cleanup(t)

			// Make request
			result := env.queryHelper(t, tc.q, nil)

			// Verify snapshot
			env.snap.Snapshot(desc, &result)
		})
	}
}
