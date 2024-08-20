package server

import (
	"apikeyper/internal/database"
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *CreateWorkspaceRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.Name == "" {
		problems["name"] = "workspace name is required"
	}

	return
}

func (s *Server) CreateWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	decodedJson, problems, err := decodeValid[*CreateWorkspaceRequest](r)

	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	var workspace = &database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: decodedJson.Name,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, createErr := s.Db.CreateWorkspace(workspace)

	if createErr != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to create workspace")
		return
	}

	respBody := CreateWorkspaceResponse{
		WorkspaceId: workspace.ID,
	}

	encode(w, r, http.StatusCreated, respBody)
}
