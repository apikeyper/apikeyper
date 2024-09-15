package database

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func (s *service) CreateWorkspace(workspace *Workspace) (uuid.UUID, error) {
	result := s.db.Create(workspace)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to create workspace. Error: %v", result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Created workspace: %v", workspace.ID))
	return workspace.ID, nil
}

func (s *service) FetchWorkspaceById(workspaceId uuid.UUID) (*Workspace, error) {
	var workspace Workspace
	result := s.db.Where("id = ?", workspaceId).First(&workspace)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch workspace with id: %s. Error: %v", workspaceId, result.Error))
		return nil, result.Error
	}
	slog.Info(fmt.Sprintf("Fetched workspace: %v", workspace.ID))
	return &workspace, nil
}

func (s *service) FetchWorkspaceByUser(workspaceUserID string) (*Workspace, error) {
	var workspace Workspace
	result := s.db.Where("workspace_user_id = ?", workspaceUserID).First(&workspace)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch workspace with id: %s. Error: %v", workspaceUserID, result.Error))
		return nil, result.Error
	}
	slog.Info(fmt.Sprintf("Fetched workspace: %v", workspace.ID))
	return &workspace, nil
}
