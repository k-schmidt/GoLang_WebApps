package example

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	res := httptest.NewRecorder()
	path := "http://localhost:4000/test"
	o, err := http.NewRequest("GET", path, nil)
	http.DefaultServeMux.ServeHTTP(res, req)
	response, err := ioutil.ReadAll(res.Body)
	if string(response) != "115" || err != nil {
		t.Errorf("Expected [], go %s", string(response))
	}
}
