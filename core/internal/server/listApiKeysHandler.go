package server

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) ListApiKeysForApiHandler(w http.ResponseWriter, r *http.Request) {

	apiId := strings.Split(r.URL.Path, "/")[2]

	if apiId == "" {
		encode(w, r, http.StatusBadRequest, "apiId is required")
		return
	}

	apiUuid, err := uuid.Parse(apiId)

	if err != nil {
		encode(w, r, http.StatusBadRequest, "Invalid apiId")
		return
	}

	apiKeys, err := s.Db.ListApiKeysForApi(apiUuid)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var respBody []FetchApiKeyResponse

	for _, ak := range *apiKeys {
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
