package apikey_test

import (
	"keyify/internal/database"
	"keyify/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateApiKey(t *testing.T) {
	// Create a new service
	s := &server.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(http.HandlerFunc(s.CreateApiKeyHandler))

	defer server.Close()

	resp, err := http.Post(server.URL, "application/json", nil)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	t.Log(resp)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

}
