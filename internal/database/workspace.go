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
