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

	if r.UserGithubId == "" {
		problems["userGithubId"] = "user github id is required"
		return
	}

	return
}

func (s *Server) FetchOrCreateWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	decodedJson, problems, err := decodeValid[*CreateWorkspaceRequest](r)

	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	// Fetch the user
	user, err := s.Db.FetchUserByGithubId(decodedJson.UserGithubId)

	if err != nil {
		encode(w, r, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	if user.Workspaces != nil && len(user.Workspaces) > 0 {
		existingWorkspace := user.Workspaces[0]
		resp := CreateWorkspaceResponse{
			WorkspaceId: existingWorkspace.ID,
		}
		encode(w, r, http.StatusOK, resp)
		return
	}

	var workspace = &database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: decodedJson.Name,
		Users:         []*database.User{user},
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
