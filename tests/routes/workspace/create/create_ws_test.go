package tests

import (
	"bytes"
	"encoding/json"
	"keyify/internal/database"
	KeyifyServer "keyify/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"keyify/tests"

	"github.com/stretchr/testify/assert"
)

func TestCreateWorkspaceHandler(t *testing.T) {
	// Create a new service
	s := &KeyifyServer.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(http.HandlerFunc(s.CreateWorkspaceHandler))

	defer server.Close()

	createRootKeyReq := KeyifyServer.CreateWorkspaceRequest{
		Name: "test-ws",
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(createRootKeyReq)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL, &buf)

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Cleanup db
	defer tests.CleanupDb()
}
