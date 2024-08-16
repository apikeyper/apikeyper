package database

import "github.com/google/uuid"

func (s *service) CreateWorkspace(workspace *Workspace) uuid.UUID {
	s.db.Create(workspace)
	return workspace.ID
}
