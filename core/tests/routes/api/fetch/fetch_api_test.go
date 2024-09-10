package tests

import (
	"apikeyper/internal/database"
	"apikeyper/internal/database/utils"
	ApikeyperServer "apikeyper/internal/server"
	"apikeyper/tests"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFetchApiByIdHandler(t *testing.T) {
	// Create a new service
	s := &ApikeyperServer.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			ApikeyperServer.Auth(
				s.Db, s.FetchApiHandler,
			),
		),
	)

	defer server.Close()

	// Create a workspace
	workspaceId, _ := s.Db.CreateWorkspace(&database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: "test-workspace",
	})

	// Create root key
	rootKey := "test-root-key"
	s.Db.CreateRootKey(&database.RootKey{
		ID:            uuid.New(),
		WorkspaceId:   workspaceId,
		RootHashedKey: utils.HashString(rootKey),
	})

	// Create an API
	apiId := uuid.New()
	s.Db.CreateApi(&database.Api{
		ID:          apiId,
		WorkspaceId: workspaceId,
	})

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/%s", server.URL, apiId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rootKey))
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
