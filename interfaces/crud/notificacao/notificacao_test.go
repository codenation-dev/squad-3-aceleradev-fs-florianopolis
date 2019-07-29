package notificacao

import (
	"testing"
)

func TestInsertNotificacao(t *testing.T) {
	//needs mocking? crud
}

func TestDelete(t *testing.T) {
	//needs mocking? crud
}

func TestGet(t *testing.T) {
	//needs mocking? crud
}

/*func Test10MaioresEstadosDoBrasil(t *testing.T) {
	estados, err := os10maioresEstadosDoBrasil()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(estados))
}

func TestQ1(t *testing.T) {
	r, err := q1()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, r)
}

func TestQuote(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote", quote())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestQuoteByActor(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/quote/{actor}", quoteByActor())
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/v1/quote/John+Cleese")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
func TestGithubStars(t *testing.T) {
	err := githubStars("go")
	assert.Nil(t, err)
	assert.FileExists(t, "stars.json")
}*/
