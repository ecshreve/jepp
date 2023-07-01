package tests

import (
	"fmt"
	"testing"
)

func TestBasicQueries(t *testing.T) {
	// Setup test environment
	env := SetupTestEnv(t)
	defer env.snap.Verify()
	defer env.Cleanup(t)

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
					edges {
						node {
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
					edges {
						node {
							id
							name
							clues {
								id
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
				clues {
					edges {
						node {
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
			// Make request
			result := env.queryHelper(t, tc.q, nil)

			// Verify snapshot
			env.snap.Snapshot(desc, &result)
		})
	}
}

func TestComplexQueries(t *testing.T) {
	// Setup test environment
	env := SetupTestEnv(t)
	defer env.snap.Verify()
	defer env.Cleanup(t)

	testcases := []struct {
		q string
	}{
		{
			q: `query {
				categories {
					edges {
						node {
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
			// Make request
			result := env.queryHelper(t, tc.q, nil)

			// Verify snapshot
			env.snap.Snapshot(desc, &result)
		})
	}
}
