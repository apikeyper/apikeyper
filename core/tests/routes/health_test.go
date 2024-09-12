package tests

import (
	"apikeyper/internal/database"
	"apikeyper/internal/server"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	s := &server.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(http.HandlerFunc(s.HealthHandler))

	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"idle\":\"1\",\"in_use\":\"0\",\"max_idle_closed\":\"0\",\"max_lifetime_closed\":\"0\",\"message\":\"It's healthy\",\"open_connections\":\"1\",\"status\":\"up\",\"wait_count\":\"0\",\"wait_duration\":\"0s\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
