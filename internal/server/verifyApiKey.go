package server

import (
	"context"
	"keyify/internal/database/utils"
	"net/http"

	"github.com/google/uuid"
)

func (r *VerifyApiKeyRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.ApiKey == "" {
		problems["apiKey"] = "apiKey is required"
	}

	if r.ApiId == uuid.Nil {
		problems["apiId"] = "apiId is required"
	}

	return
}

func (s *Server) VerifyApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	decodedJson, problems, err := decodeValid[*VerifyApiKeyRequest](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	// Fetch the api
	haskedKey := utils.HashString(decodedJson.ApiKey)
	api, err := s.Db.VerifyApiKey(haskedKey)
	if err != nil {
		encode(w, r, http.StatusNotFound, "Failed to verify api key")
		return
	}

	respBody := VerifyApiKeyResponse{
		ApiId: api.ID,
		KeyId: api.ID,
		Valid: true,
	}

	encode(w, r, http.StatusCreated, respBody)
}
