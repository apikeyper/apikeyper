package server

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) FetchApiHandler(w http.ResponseWriter, r *http.Request) {

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

	api, err := s.Db.FetchApi(apiUuid)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve apis for workspace")
		return
	}

	respBody := FetchApiResponse{
		ApiId:       api.ID,
		WorkspaceId: api.WorkspaceId,
		ApiName:     api.ApiName,
		CreatedAt:   api.CreatedAt,
		UpdatedAt:   api.UpdatedAt,
	}

	encode(w, r, http.StatusOK, respBody)
}
