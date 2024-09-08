package tests

import (
	"apikeyper/internal/database"
	ApikeyperServer "apikeyper/internal/server"
	"apikeyper/tests"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFetchWorkspaceByIdHandler(t *testing.T) {
	// Create a new service
	s := &ApikeyperServer.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			s.FetchWorkspaceHandler,
		),
	)

	defer server.Close()

	// Create a workspace
	workspaceId, _ := s.Db.CreateWorkspace(&database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: "test-workspace",
	})

	client := &http.Client{}
	url := fmt.Sprintf("%s/workspace/%s", server.URL, workspaceId)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Cleanup db
	defer tests.CleanupDb()
}
