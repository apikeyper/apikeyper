package server

import (
	"encoding/json"
	"keyify/internal/database"
	"keyify/internal/database/utils"
	"keyify/internal/schemas"
	"log"
	"net/http"
	"time"
)

func (s *Server) CreateApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	generatedApiKey, err := utils.GenerateApiKey("keyify_")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error generating api key"))
		return
	}

	keyId := utils.GenerateRandomId("key_")

	hashedKey := utils.HashString(generatedApiKey)

	var apiKeyRow = &database.ApiKey{
		KeyId:       keyId,
		ApiId:       "api1",
		WorkspaceId: "ws1",
		HashedKey:   hashedKey,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	log.Print("1")

	log.Print(apiKeyRow)

	s.Db.CreateApiKey(apiKeyRow)

	log.Print("2")

	w.WriteHeader(http.StatusCreated)

	respBody := schemas.CreateKeyResponse{
		ApiId: apiKeyRow.ApiId,
		KeyId: keyId,
		Key:   generatedApiKey,
	}

	// Marshal the response body
	respBodyJSON, err := json.Marshal(respBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error marshalling response"))
		return
	}

	_, _ = w.Write(respBodyJSON)
}
