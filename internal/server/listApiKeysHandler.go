package server

import (
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) ListApiKeysForApiHandler(w http.ResponseWriter, r *http.Request) {

	apiId := r.URL.Query().Get("apiId")

	apiUuid, err := uuid.Parse(apiId)

	apiKeyUsageRecords, err := s.Db.ListApiKeysForApi(apiUuid)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve keys for api")
		return
	}

	var respBody []FetchApiKeyResponse

	for _, ak := range *apiKeyUsageRecords {
		respBody = append(respBody, FetchApiKeyResponse{
			ApiKeyId:  ak.ID,
			ApiId:     ak.ApiId,
			Name:      ak.Name,
			ExpiresAt: ak.ExpiresAt,
			Status:    ak.Status,
			CreatedAt: ak.CreatedAt,
			UpdatedAt: ak.UpdatedAt,
		})
	}

	encode(w, r, http.StatusOK, respBody)
}
