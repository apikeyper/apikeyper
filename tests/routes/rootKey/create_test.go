package tests

import (
	"bytes"
	"encoding/json"
	"keyify/internal/database"
	KeyifyServer "keyify/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateRootKeyHandler(t *testing.T) {
	// Create a new service
	s := &KeyifyServer.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(http.HandlerFunc(s.CreateRootKeyHandler))

	defer server.Close()

	createRootKeyReq := KeyifyServer.CreateRootKeyRequest{
		Name:        "test-root-key",
		WorkspaceId: uuid.New(),
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(createRootKeyReq)

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.URL, &buf)

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
