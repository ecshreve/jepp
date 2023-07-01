package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/graph/common"
	graph "github.com/ecshreve/jepp/graph/generated"
	"github.com/ecshreve/jepp/graph/model"
	resolvers "github.com/ecshreve/jepp/graph/resolvers"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestEnv struct {
	db         *gorm.DB
	gqlserver  *handler.Server
	httpserver *httptest.Server
	clk        *clock.Mock

	snap *snapshotter.Snapshotter
}

// Setup the test environment.
func SetupTestEnv(t *testing.T) *TestEnv {
	env := &TestEnv{}

	// Setup Clock
	clk := clock.NewMock()
	clk.Set(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))
	env.clk = clk

	// Setup Database
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	require.NoError(t, err)
	err = db.AutoMigrate(&model.Clue{}, &model.Category{}, &model.Game{}, &model.Season{})
	require.NoError(t, err)
	env.db = db

	// Setup GraphQL Server
	os.Setenv("PORT", "4000")
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))
	require.NotNil(t, srv)
	env.gqlserver = srv

	// Setup HTTP Server
	customContext := common.CreateContext(&common.CustomContext{Database: db}, srv)
	httpserver := httptest.NewServer(customContext)
	require.NotNil(t, httpserver)
	env.httpserver = httpserver

	// Setup Snapshotter
	snap := snapshotter.New(t)
	snap.SnapshotErrors = true
	env.snap = snap

	// Seed Data
	env.SeedData(t)

	return env
}

// Run After All Test Done
func (e *TestEnv) Cleanup(t *testing.T) {
	// Drop Tables
	for _, val := range []interface{}{&model.Clue{}, &model.Category{}, &model.Game{}, &model.Season{}} {
		e.db.Migrator().DropTable(val)
	}
}

func (e *TestEnv) SeedData(t *testing.T) {
	seasons := []model.Season{
		{ID: 1, Number: 1, StartDate: e.clk.Now(), EndDate: e.clk.Now().Add(time.Hour * 24 * 30)},
		{ID: 2, Number: 2, StartDate: e.clk.Now().Add(time.Hour * 24 * 31), EndDate: e.clk.Now().Add(time.Hour * 24 * 60)},
		{ID: 3, Number: 3, StartDate: e.clk.Now().Add(time.Hour * 24 * 61), EndDate: e.clk.Now().Add(time.Hour * 24 * 90)},
	}

	games := []model.Game{
		{ID: 1, SeasonID: 1, Show: 1, AirDate: e.clk.Now().Add(time.Hour * 24 * 2), TapeDate: e.clk.Now()},
		{ID: 2, SeasonID: 1, Show: 2, AirDate: e.clk.Now().Add(time.Hour * 24 * 4), TapeDate: e.clk.Now().Add(time.Hour * 24 * 3)},
		{ID: 3, SeasonID: 1, Show: 3, AirDate: e.clk.Now().Add(time.Hour * 24 * 6), TapeDate: e.clk.Now().Add(time.Hour * 24 * 5)},
	}

	categories := []model.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
		{ID: 3, Name: "Category 3"},
	}

	clues := []model.Clue{
		{ID: 1, CategoryID: 1, GameID: 1, Question: "Question 1", Answer: "Answer 1"},
		{ID: 2, CategoryID: 1, GameID: 1, Question: "Question 2", Answer: "Answer 2"},
		{ID: 3, CategoryID: 1, GameID: 1, Question: "Question 3", Answer: "Answer 3"},
		{ID: 4, CategoryID: 2, GameID: 1, Question: "Question 4", Answer: "Answer 4"},
		{ID: 5, CategoryID: 2, GameID: 1, Question: "Question 5", Answer: "Answer 5"},
		{ID: 6, CategoryID: 2, GameID: 1, Question: "Question 6", Answer: "Answer 6"},
		{ID: 7, CategoryID: 3, GameID: 1, Question: "Question 7", Answer: "Answer 7"},
		{ID: 8, CategoryID: 3, GameID: 1, Question: "Question 8", Answer: "Answer 8"},
		{ID: 9, CategoryID: 3, GameID: 1, Question: "Question 9", Answer: "Answer 9"},
		{ID: 10, CategoryID: 1, GameID: 2, Question: "Question 10", Answer: "Answer 10"},
		{ID: 11, CategoryID: 1, GameID: 2, Question: "Question 11", Answer: "Answer 11"},
		{ID: 12, CategoryID: 1, GameID: 2, Question: "Question 12", Answer: "Answer 12"},
		{ID: 13, CategoryID: 2, GameID: 2, Question: "Question 13", Answer: "Answer 13"},
		{ID: 14, CategoryID: 2, GameID: 2, Question: "Question 14", Answer: "Answer 14"},
		{ID: 15, CategoryID: 2, GameID: 2, Question: "Question 15", Answer: "Answer 15"},
		{ID: 16, CategoryID: 3, GameID: 2, Question: "Question 16", Answer: "Answer 16"},
		{ID: 17, CategoryID: 3, GameID: 2, Question: "Question 17", Answer: "Answer 17"},
		{ID: 18, CategoryID: 3, GameID: 2, Question: "Question 18", Answer: "Answer 18"},
		{ID: 19, CategoryID: 1, GameID: 3, Question: "Question 19", Answer: "Answer 19"},
		{ID: 20, CategoryID: 1, GameID: 3, Question: "Question 20", Answer: "Answer 20"},
	}

	e.db.Create(seasons)
	e.db.Create(games)
	e.db.Create(categories)
	e.db.Create(clues)
}

// queryHelper is a helper function to run a query and return the response
func (env *TestEnv) queryHelper(t *testing.T, query string, variables map[string]interface{}) map[string]interface{} {
	// Convert query string to JSON
	kv := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	postJson, err := json.Marshal(kv)
	assert.NoError(t, err)

	// Send request
	resp, err := http.Post(env.httpserver.URL, "application/json", bytes.NewBuffer(postJson))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t, err)

	// Unmarshal response body
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	return result
}
