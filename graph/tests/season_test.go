package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ecshreve/jepp/app/models"
	"github.com/kr/pretty"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Results struct {
	Data struct {
		Seasons []*models.Season `json:"seasons"`
	} `json:"data"`
}

func (t *SuiteTest) TestSeason() {
	snap := snapshotter.New(t.T())
	snap.SnapshotErrors = true
	defer snap.Verify()

	q := `query {
		seasons {
			id
			number
			startDate
			endDate
		}
	}`

	kv := map[string]string{
		"query": q,
	}
	postJson, err := json.Marshal(kv)
	assert.NoError(t.T(), err)

	resp, err := http.Post(t.httpserver.URL, "application/json", bytes.NewBuffer(postJson))
	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), resp)
	assert.Equal(t.T(), 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	assert.NoError(t.T(), err)
	pretty.Print(string(body))

	var result Results
	err = json.Unmarshal(body, &result)
	pretty.Print(result)
	require.NoError(t.T(), err)
	snap.Snapshot("seasons", &result)
}

// 	req := httptest.NewRequest("POST", "http://localhost/query", bytes.NewBufferString(query))
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	resp := w.Result()
// 	body, _ := io.ReadAll(resp.Body)

// 	fmt.Println(resp.StatusCode)
// 	fmt.Println(resp.Header.Get("Content-Type"))
// 	fmt.Println(string(body))
// }
