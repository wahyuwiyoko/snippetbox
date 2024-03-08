package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	result, err := ts.Client().Get(ts.URL + urlPath)

	if err != nil {
		t.Fatal(err)
	}

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)

	if err != nil {
		t.Fatal(err)
	}

	return result.StatusCode, result.Header, body
}
