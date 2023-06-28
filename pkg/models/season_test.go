package models_test

import (
	"testing"

	"github.com/ecshreve/jepp/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSeasons(t *testing.T) {
	testdb := models.GetDBHandle()
	defer testdb.Close()

	seasons, err := models.GetSeasons()
	assert.NoError(t, err)
	assert.NotEmpty(t, seasons)
	assert.Equal(t, 39, len(seasons))
}
