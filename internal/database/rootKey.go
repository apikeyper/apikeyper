package database

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
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

func (s *service) FetchRootKeyById(rootKeyId uuid.UUID) (*RootKey, error) {
	var rootKey RootKey
	result := s.db.Where("id = ?", rootKeyId).First(&rootKey)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch root key with id: %s. Error: %v", rootKeyId, result.Error))
		return nil, result.Error
	}
	slog.Info(fmt.Sprintf("Fetched root key: %v", rootKey.ID))
	return &rootKey, nil
}

func (s *service) ListRootKeysForWorkspace(workspaceId uuid.UUID) (*[]RootKey, error) {
	var rootKeys []RootKey
	result := s.db.
		Order(
			clause.OrderByColumn{
				Column: clause.Column{Name: "created_at"},
				Desc:   true,
			},
		).
		Where("workspace_id = ?", workspaceId).
		Find(&rootKeys)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch root keys for workspace: %s. Error: %v", workspaceId, result.Error))
		return nil, result.Error
	}

	return &rootKeys, nil
}
