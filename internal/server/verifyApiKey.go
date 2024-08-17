package server

import (
	"context"
	"keyify/internal/database/utils"
	"keyify/internal/events"
	"net/http"
	"time"

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
	// Get session from context
	session := r.Context().Value("session").(Session)

	// Validate the request body
	decodedJson, problems, err := decodeValid[*VerifyApiKeyRequest](r)
	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	// Fetch the api key
	haskedKey := utils.HashString(decodedJson.ApiKey)
	apiKey, err := s.Db.VerifyApiKey(haskedKey)
	if err != nil {
		encode(w, r, http.StatusNotFound, "Failed to verify api key")
		return
	}

	// Publish an event for successful verification
	go s.Message.Publish(context.Background(), events.EventPayload{
		EventType: events.API_KEY_VERIFY_SUCCESS,
		Data: events.EventData{
			WorkspaceId: session.WorkspaceId.String(),
			ApiKeyId:    apiKey.ID.String(),
			ApiId:       apiKey.ApiId.String(),
			EventTime:   time.Now().String(),
		},
	})

	respBody := VerifyApiKeyResponse{
		ApiId: apiKey.ID,
		KeyId: apiKey.ID,
		Valid: true,
	}

	encode(w, r, http.StatusCreated, respBody)
}
