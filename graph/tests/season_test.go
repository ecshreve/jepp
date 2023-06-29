package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func (t *SuiteTest) TestSeason() {
	q := `query GetSeasons {
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

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
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
