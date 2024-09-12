package tests

import (
	"apikeyper/internal/database"
	"apikeyper/internal/database/utils"
	"apikeyper/internal/events"
	"apikeyper/internal/ratelimit"
	ApikeyperServer "apikeyper/internal/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"apikeyper/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListApiKeysHandler(t *testing.T) {
	// Create a new service
	s := &ApikeyperServer.Server{
		Db:          database.New(),
		Message:     events.New(),
		RateLimiter: ratelimit.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			ApikeyperServer.Auth(
				s.Db, s.ListApiKeysForApiHandler,
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

	// Create an API keys
	generatedApiKey1, _ := utils.GenerateApiKey("test_")
	prefix1 := "test_"
	name1 := "test-api-key"
	s.Db.CreateApiKey(&database.ApiKey{
		ID:        uuid.New(),
		ApiId:     apiId,
		HashedKey: utils.HashString(generatedApiKey1),
		Name:      &name1,
		Prefix:    &prefix1,
	})

	generatedApiKey2, _ := utils.GenerateApiKey("test_")
	prefix2 := "test2_"
	name2 := "test-api-key-2"
	s.Db.CreateApiKey(&database.ApiKey{
		ID:        uuid.New(),
		ApiId:     apiId,
		HashedKey: utils.HashString(generatedApiKey2),
		Name:      &name2,
		Prefix:    &prefix2,
	})

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/%s/keys", server.URL, apiId)
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
