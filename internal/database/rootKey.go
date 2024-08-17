package database

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func (s *service) CreateRootKey(rootKey *RootKey) (uuid.UUID, error) {
	result := s.db.Create(rootKey)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to create root key. Error: %v", result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Created root key: %v for workspaceId: %v", rootKey.ID, rootKey.WorkspaceId))
	return rootKey.ID, nil
}

func (s *service) FetchRootKey(rootHashedKey string) (*RootKey, error) {
	var rootKey RootKey
	if result := s.db.Where("root_hashed_key = ?", rootHashedKey).First(&rootKey); result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch root key. Error: %v", result.Error))
		return nil, result.Error
	}

	return &rootKey, nil
}
