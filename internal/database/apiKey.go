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

	slog.Info(fmt.Sprintf("Fetched api key to verify: %v", apiKey.ID))
	return &apiKey, nil
}

func (s *service) UpdateApiKeyStatus(apiKeyId uuid.UUID, status string) (*ApiKey, error) {
	var apiKey ApiKey
	result := s.db.Where("id = ?", apiKeyId).First(&apiKey)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to update api key status for api key: %s. Error: %v", apiKeyId, result.Error))
		return nil, result.Error
	}

	apiKey.Status = status

	result = s.db.Save(&apiKey)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to update api key status for api key: %s. Error: %v", apiKeyId, result.Error))
		return nil, result.Error
	}

	slog.Info(fmt.Sprintf("Updated api key status: %v for api key: %s", status, apiKeyId))
	return &apiKey, nil
}

func (s *service) LogApiKeyUsage(apiKeyUsage *ApiKeyUsage) (uuid.UUID, error) {
	result := s.db.Create(apiKeyUsage)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to log api key usage for api key: %s. Error: %v", apiKeyUsage.ApiKeyId, result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Logged api key usage: %v for api key: %s", apiKeyUsage.ID, apiKeyUsage.ApiKeyId))
	return apiKeyUsage.ID, nil
}
