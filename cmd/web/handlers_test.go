package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	ts := httptest.NewTLSServer(app.routes())

	defer ts.Close()

	result, err := ts.Client().Get(ts.URL + "/ping")

	if err != nil {
		t.Fatal(err)
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, result.StatusCode)
	}

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
