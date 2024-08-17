package server

import (
	"context"
	"keyify/internal/database"
	"keyify/internal/database/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *CreateApiKeyRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.ApiId == uuid.Nil {
		problems["apiId"] = "apiId is required"
	}

	return
}

func determinePrefix(r *CreateApiKeyRequest) string {
	if r.Prefix != "" {
		return r.Prefix
	} else {
		return "keyify_"
	}
}
func (s *Server) CreateApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	decodedJson, problems, err := decodeValid[*CreateApiKeyRequest](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	// Fetch the api
	api, err := s.Db.FetchApi(decodedJson.ApiId)
	if err != nil {
		encode(w, r, http.StatusNotFound, "Api not found")
		return
	}

	// Generate the api key
	keyPrefix := determinePrefix(decodedJson)
	generatedApiKey, err := utils.GenerateApiKey(keyPrefix)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Error generating api key")
		return
	}

	hashedKey := utils.HashString(generatedApiKey)

	var apiKeyRow = &database.ApiKey{
		ID:        uuid.New(),
		ApiId:     api.ID,
		Name:      &decodedJson.Name,
		Prefix:    &decodedJson.Prefix,
		HashedKey: hashedKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, createErr := s.Db.CreateApiKey(apiKeyRow)

	if createErr != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to create api key")
		return
	}

	respBody := CreateKeyResponse{
		ApiId: api.ID,
		KeyId: apiKeyRow.ID,
		Key:   generatedApiKey,
	}

	encode(w, r, http.StatusCreated, respBody)
}
