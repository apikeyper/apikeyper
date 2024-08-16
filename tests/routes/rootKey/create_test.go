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

	// Create a workspace
	workspaceId := s.Db.CreateWorkspace(&database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: "test-workspace",
	})

	createRootKeyReq := KeyifyServer.CreateRootKeyRequest{
		Name:        "test-root-key",
		WorkspaceId: workspaceId,
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
