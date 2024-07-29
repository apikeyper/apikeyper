package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keyify/internal/database"
	"keyify/internal/database/utils"
	KeyifyServer "keyify/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateApiHandler(t *testing.T) {
	// Create a new service
	s := &KeyifyServer.Server{
		Db: database.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			KeyifyServer.Auth(
				s.Db, s.CreateApiHandler,
			),
		),
	)

	defer server.Close()

	// Create root key
	rootKey := "test-root-key"
	s.Db.CreateRootKey(&database.RootKey{
		ID:            uuid.New(),
		RootHashedKey: utils.HashString(rootKey),
	})

	createApiReq := KeyifyServer.CreateApiRequest{
		ApiName: "test-api",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(createApiReq)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL, &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rootKey))

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
