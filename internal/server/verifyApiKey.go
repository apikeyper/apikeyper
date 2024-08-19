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

func SendApiKeyVerifySuccessEvent(
	messageService events.MessageService,
	workspaceId uuid.UUID,
	apiKeyId uuid.UUID,
	apiId uuid.UUID,
	eventTime time.Time,
) {
	go messageService.Publish(context.Background(), events.EventPayload{
		EventType: events.API_KEY_VERIFY_SUCCESS,
		Data: events.EventData{
			EventId:     uuid.New(),
			WorkspaceId: workspaceId.String(),
			ApiKeyId:    apiKeyId.String(),
			ApiId:       apiId.String(),
			EventTime:   eventTime.String(),
		},
	})
}

func SendApiKeyVerifyFailedEvent(
	messageService events.MessageService,
	workspaceId uuid.UUID,
	apiKeyId uuid.UUID,
	apiId uuid.UUID,
	eventTime time.Time,
) {
	go messageService.Publish(context.Background(), events.EventPayload{
		EventType: events.API_KEY_VERIFY_FAILED,
		Data: events.EventData{
			EventId:     uuid.New(),
			WorkspaceId: workspaceId.String(),
			ApiKeyId:    apiKeyId.String(),
			ApiId:       apiId.String(),
			EventTime:   eventTime.String(),
		},
	})
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

	// Check if the api key is active
	if apiKey.Status != "active" {
		encode(w, r, http.StatusForbidden, VerifyApiKeyResponse{
			ApiId: apiKey.ApiId,
			KeyId: apiKey.ID,
			Valid: false,
		})
		SendApiKeyVerifyFailedEvent(s.Message, session.WorkspaceId, apiKey.ID, apiKey.ApiId, time.Now())
		return
	}

	// Check if the api key is expired
	if apiKey.ExpiresAt.Valid && apiKey.ExpiresAt.Time.Before(time.Now()) {
		encode(w, r, http.StatusForbidden, VerifyApiKeyResponse{
			ApiId: apiKey.ApiId,
			KeyId: apiKey.ID,
			Valid: false,
		})
		s.Db.UpdateApiKeyStatus(apiKey.ID, "expired")
		SendApiKeyVerifyFailedEvent(s.Message, session.WorkspaceId, apiKey.ID, apiKey.ApiId, time.Now())
		return
	}

	// Publish an event for successful verification
	SendApiKeyVerifySuccessEvent(s.Message, session.WorkspaceId, apiKey.ID, apiKey.ApiId, time.Now())

	respBody := VerifyApiKeyResponse{
		ApiId: apiKey.ApiId,
		KeyId: apiKey.ID,
		Valid: true,
	}

	encode(w, r, http.StatusCreated, respBody)
}
