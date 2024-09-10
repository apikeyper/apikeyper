package server

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) FetchRootKeyByIdHandler(w http.ResponseWriter, r *http.Request) {

	rootKeyId := strings.Split(r.URL.Path, "/")[2]

	if rootKeyId == "" {
		encode(w, r, http.StatusBadRequest, "rootKeyId is required")
		return
	}

	rootKeyUuid, err := uuid.Parse(rootKeyId)

	if err != nil {
		encode(w, r, http.StatusBadRequest, "Invalid rootKeyId")
		return
	}

	rootKey, err := s.Db.FetchRootKeyById(rootKeyUuid)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve rootKey for workspace")
		return
	}

	respBody := FetchRootKeyResponse{
		RootKeyId:   rootKey.ID,
		WorkspaceId: rootKey.WorkspaceId,
		RootKeyName: rootKey.RootKeyName,
		CreatedAt:   rootKey.CreatedAt,
		UpdatedAt:   rootKey.UpdatedAt,
	}

	encode(w, r, http.StatusOK, respBody)
}
