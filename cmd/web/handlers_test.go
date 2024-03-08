package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())

	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old silent pond...")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.get(t, test.urlPath)

			if code != test.wantCode {
				t.Errorf("want %d; got %d", test.wantCode, code)
			}

			if !bytes.Contains(body, test.wantBody) {
				t.Errorf("want body to contain %q", test.wantBody)
			}
		})
	}
}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	const (
		validName     = "Bob"
		validPassword = "validPa$$word"
		validEmail    = "bob@example.com"
	)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", validName, validEmail, validPassword, csrfToken, http.StatusSeeOther, nil},
		{"Empty name", "", validEmail, validPassword, csrfToken, http.StatusOK, []byte("This field cannot be empty")},
		{"Empty email", validName, "", validPassword, csrfToken, http.StatusOK, []byte("This field cannot be empty")},
		{"Empty password", validName, validEmail, "", csrfToken, http.StatusOK, []byte("This field cannot be empty")},
		{"Invalid email (incomplete domain)", validName, "bob@example.", validPassword, csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", validName, "bobexample.com", validPassword, csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing local part)", validName, "@example.com", validPassword, csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", validName, validEmail, "pa$$word", csrfToken, http.StatusOK, []byte("This field is too short (minimum is 10 characters)")},
		{"Duplicate email", validName, "dupe@example.com", validPassword, csrfToken, http.StatusOK, []byte("Address is already in use!")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			form := url.Values{}

			form.Add("name", test.userName)
			form.Add("email", test.userEmail)
			form.Add("password", test.userPassword)
			form.Add("csrf_token", test.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			if code != test.wantCode {
				t.Errorf("want %d; got %d", test.wantCode, code)
			}

			if !bytes.Contains(body, test.wantBody) {
				t.Errorf("want body %s to contain %q", body, test.wantBody)
			}
		})
	}
}
