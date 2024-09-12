package server

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (s *Server) FetchWorkspaceHandler(w http.ResponseWriter, r *http.Request) {

	workspaceId := strings.Split(r.URL.Path, "/")[2]

	if workspaceId == "" {
		encode(w, r, http.StatusBadRequest, "workspaceId is required")
		return
	}

	workspaceUuid, err := uuid.Parse(workspaceId)

	if err != nil {
		encode(w, r, http.StatusBadRequest, "Invalid workspaceId")
		return
	}

	workspace, err := s.Db.FetchWorkspaceById(workspaceUuid)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to retrieve workspace")
		return
	}

	respBody := FetchWorkspaceResponse{
		WorkspaceId:   workspace.ID,
		WorkspaceName: workspace.WorkspaceName,
		CreatedAt:     workspace.CreatedAt,
		UpdatedAt:     workspace.UpdatedAt,
	}

	encode(w, r, http.StatusOK, respBody)
}
