package tests

import (
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/benbjohnson/clock"
	"github.com/ecshreve/jepp/app/models"
	"github.com/ecshreve/jepp/graph"
	"github.com/ecshreve/jepp/graph/common"
	resolvers "github.com/ecshreve/jepp/graph/resolvers"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SuiteTest struct {
	suite.Suite
	db         *gorm.DB
	gqlserver  *handler.Server
	httpserver *httptest.Server
	clk        *clock.Mock
}

func TestSuite(t *testing.T) {
	os.Setenv("PORT", "4000")
	suite.Run(t, new(SuiteTest))
}

// Setup db value
func (t *SuiteTest) SetupSuite() {
	clk := clock.NewMock()
	clk.Set(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))
	t.clk = clk

	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t.T(), err)

	// db.AutoMigrate(&models.Clue{}, &models.Category{}, &models.Game{}, &models.Season{})
	db.AutoMigrate(&models.Season{})

	seasons := []models.Season{
		{ID: 1, Number: 1, StartDate: clk.Now(), EndDate: clk.Now().Add(time.Hour * 24 * 30)},
		{ID: 2, Number: 2, StartDate: clk.Now().Add(time.Hour * 24 * 31), EndDate: clk.Now().Add(time.Hour * 24 * 60)},
		{ID: 3, Number: 3, StartDate: clk.Now().Add(time.Hour * 24 * 61), EndDate: clk.Now().Add(time.Hour * 24 * 90)},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))
	t.db = db
	t.gqlserver = srv

	customContext := common.CreateContext(&common.CustomContext{Database: db}, srv)
	t.httpserver = httptest.NewServer(customContext)

	t.db.Create(seasons)
}

// Run After All Test Done
func (t *SuiteTest) TearDownSuite() {
	// Drop Table
	for _, val := range []interface{}{&models.Clue{}, &models.Category{}, &models.Game{}, &models.Season{}} {
		t.db.Migrator().DropTable(val)
	}
}

// Run Before a Test
func (t *SuiteTest) SetupTest() {

}

// Run After a Test
func (t *SuiteTest) TearDownTest() {

}
