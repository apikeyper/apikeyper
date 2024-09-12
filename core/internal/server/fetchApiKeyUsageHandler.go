package server

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) FetchApiKeyUsageHandler(w http.ResponseWriter, r *http.Request) {

	apiKeyId := strings.Split(r.URL.Path, "/")[2]
	interval := r.URL.Query().Get("interval")

	if apiKeyId == "" {
		encode(w, r, http.StatusBadRequest, "apiKeyId is required")
		return
	}

	apiKeyUuid, err := uuid.Parse(apiKeyId)

	if err != nil {
		encode(w, r, http.StatusBadRequest, "Invalid apiKeyId")
		return
	}

	apiKeyUsageRecords, err := s.Db.FetchApiKeyUsage(apiKeyUuid, interval)
	count := len(*apiKeyUsageRecords)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve apis for workspace")
		return
	}

	respBody := FetchApiKeyUsageResponse{
		ApiKeyId: apiKeyUuid,
		Records:  *apiKeyUsageRecords,
		Count:    count,
	}

	encode(w, r, http.StatusCreated, respBody)
}
