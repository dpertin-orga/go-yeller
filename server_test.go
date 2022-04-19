package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestYellingHandler(t *testing.T) {
	w := httptest.NewRecorder()
	word := "hello world"
	req := httptest.NewRequest(http.MethodGet, "/yelling?word="+strings.ReplaceAll(word, " ", "%20"), nil)
	yellingHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	expectedBody := new(bytes.Buffer)
	YellingFormat(word, expectedBody)

	if string(data) != expectedBody.String() {
		t.Errorf("expected %v, got %v", expectedBody.String(), (data))
	}
}
