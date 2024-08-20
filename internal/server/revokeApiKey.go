package server

import (
	"apikeyper/internal/database/utils"
	"apikeyper/internal/events"
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *RevokeApiKeyRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.ApiKey == "" {
		problems["apiKey"] = "apiKey is required"
	}

	if r.ApiId == uuid.Nil {
		problems["apiId"] = "apiId is required"
	}

	return
}

func SendApiKeyRevokedEvent(
	messageService events.MessageService,
	workspaceId uuid.UUID,
	apiKeyId uuid.UUID,
	apiId uuid.UUID,
	eventTime time.Time,
) {
	go messageService.Publish(context.Background(), events.EventPayload{
		EventType: events.API_KEY_REVOKED,
		Data: events.EventData{
			EventId:     uuid.New(),
			WorkspaceId: workspaceId.String(),
			ApiKeyId:    apiKeyId.String(),
			ApiId:       apiId.String(),
			EventTime:   eventTime.String(),
		},
	})
}

func (s *Server) RevokeApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Get session from context
	session := r.Context().Value("session").(Session)

	// Validate the request body
	decodedJson, problems, err := decodeValid[*RevokeApiKeyRequest](r)
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

	// Check if the api key is active
	if apiKey.Status == "revoked" || apiKey.Status == "expired" {
		encode(w, r, http.StatusForbidden, "Api key is already revoked or expired")
		return
	}

	// Check if the api key is expired
	if apiKey.Status == "active" && apiKey.ExpiresAt.Valid && apiKey.ExpiresAt.Time.Before(time.Now()) {
		encode(w, r, http.StatusAccepted, "Api key was already expired")
		s.Db.UpdateApiKeyStatus(apiKey.ID, "expired")
		return
	}

	// Publish an event for successful verification
	s.Db.UpdateApiKeyStatus(apiKey.ID, "revoked")
	SendApiKeyRevokedEvent(s.Message, session.WorkspaceId, apiKey.ID, apiKey.ApiId, time.Now())

	respBody := RevokeApiKeyResponse{
		ApiId:   apiKey.ApiId,
		KeyId:   apiKey.ID,
		Revoked: true,
	}

	encode(w, r, http.StatusCreated, respBody)
}
