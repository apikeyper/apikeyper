package database

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (s *service) CreateApi(api *Api) (uuid.UUID, error) {
	result := s.db.Create(api)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to created api. Error: %v", result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Created api: %v with for workspaceId:%v", api.ID, api.WorkspaceId))
	return api.ID, nil
}

func (s *service) FetchApi(apiId uuid.UUID) (*Api, error) {
	var api *Api
	result := s.db.Where("id = ?", apiId).Preload("ApiKeys").First(&api)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch api. Error: %v", result.Error))
		return nil, result.Error
	}

	return api, nil
}

func (s *service) ListApis(workspaceId uuid.UUID) (*[]Api, error) {
	var apis *[]Api
	result := s.db.
		Order(
			clause.OrderByColumn{
				Column: clause.Column{Name: "created_at"},
				Desc:   true,
			},
		).
		Where("workspace_id = ?", workspaceId).
		Find(&apis)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to list apis. Error: %v", result.Error))
		return nil, result.Error
	}

	return apis, nil
}
