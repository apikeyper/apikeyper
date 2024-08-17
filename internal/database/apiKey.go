package database

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func (s *service) CreateApiKey(apiKey *ApiKey) (uuid.UUID, error) {
	result := s.db.Create(apiKey)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to created api key for apiId: %s. Error: %v", apiKey.ApiId, result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Created api key: %v with for apiId:%v", apiKey.ID, apiKey.ApiId))
	return apiKey.ID, nil
}

func (s *service) FetchApiKeyById(apiKeyId string) (*ApiKey, error) {
	var apiKey ApiKey
	result := s.db.Where("id = ?", apiKeyId).First(&apiKey)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch api key with id: %s. Error: %v", apiKeyId, result.Error))
		return nil, result.Error
	}

	slog.Info(fmt.Sprintf("Fetched api key: %v with id: %s", apiKey.ID, apiKeyId))
	return &apiKey, nil
}

func (s *service) VerifyApiKey(apiKeyHashed string) (*ApiKey, error) {
	var apiKey ApiKey
	result := s.db.Where("hashed_key = ?", apiKeyHashed).First(&apiKey).Select("id")

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to verify api key with hashed key: %s. Error: %v", apiKeyHashed, result.Error))
		return nil, result.Error
	}

	slog.Info(fmt.Sprintf("Verified api key: %v with hashed key: %s", apiKey.ID, apiKeyHashed))
	return &apiKey, nil
}
