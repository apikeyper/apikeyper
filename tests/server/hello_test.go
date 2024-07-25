package tests

import (
	"io"
	"keyify/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Log("creating server")
	s := &server.Server{}

	server := httptest.NewServer(http.HandlerFunc(s.HealthHandler))

	defer server.Close()

	t.Log(server.URL)

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
