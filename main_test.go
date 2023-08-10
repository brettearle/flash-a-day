package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Path: main.go
func TestMain(t *testing.T) {
	//integration test for server started and end point / returns Hello, World!

	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleGetFunc)
	server := httptest.NewServer(mux)
	ts := server.Client()
	want := "Hell"

	res, err := ts.Get(server.URL)
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("Error: %q", err)
	}
	got := string(body)
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

}
