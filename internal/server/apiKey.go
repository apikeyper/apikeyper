package server

import (
	"keyify/internal/database"
	"net/http"
	"time"
)

func (s *Server) CreateApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	var apiKey = &database.ApiKey{
		KeyId:     "apiKey1",
		KeyHash:   "keyHash",
		CreatedAt: time.Now(),
	}

	apiKeyId := s.db.CreateApiKey(apiKey)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(apiKeyId))
}
