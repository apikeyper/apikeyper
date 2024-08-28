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

func (s *service) FetchApiKeyById(apiKeyId uuid.UUID) (*ApiKey, error) {
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
	result := s.db.Where("hashed_key = ?", apiKeyHashed).Preload("RateLimitConfig").First(&apiKey).Select("id")

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

func (s *service) LogApiKeyUsage(apiKeyUsage *ApiKeyActivity) (uuid.UUID, error) {
	result := s.db.Create(apiKeyUsage)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to log api key usage for api key: %s. Error: %v", apiKeyUsage.ApiKeyId, result.Error))
		return uuid.Nil, result.Error
	}

	slog.Info(fmt.Sprintf("Logged api key usage: %v for api key: %s", apiKeyUsage.ID, apiKeyUsage.ApiKeyId))
	return apiKeyUsage.ID, nil
}

func (s *service) FetchApiKeyUsage(apiKeyId uuid.UUID, interval string) (*[]ApiKeyUsageCount, error) {
	var apiKeyActivityRecords []ApiKeyUsageCount

	apiKeyIdStr := apiKeyId.String()

	stmt := fmt.Sprintf(`WITH interval_data AS (
							SELECT date_trunc('minute', created_at) - (date_part('minute', created_at)::integer %% %v || ' minutes')::interval AS interval_start,
										usage,
										COUNT(*) AS total_usage
							FROM api_key_activities
							WHERE "api_key_id" = '%v'
							GROUP BY interval_start, usage
							)
							SELECT interval_start,
										MAX(CASE WHEN usage = 'failed' THEN total_usage END) AS failed,
										MAX(CASE WHEN usage = 'revoked' THEN total_usage END) AS revoked,
										MAX(CASE WHEN usage = 'success' THEN total_usage END) AS success,
										MAX(CASE WHEN usage = 'rate_limited' THEN total_usage END) AS rate_limited
							FROM interval_data
							GROUP BY interval_start;`,
		interval,
		apiKeyIdStr,
	)

	result := s.db.Raw(stmt).Scan(&apiKeyActivityRecords)

	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch api key usage for api key: %s. Error: %v", apiKeyId, result.Error))
		return nil, result.Error
	}

	slog.Info(fmt.Sprintf("Fetched api key usage for api key: %s", apiKeyId))

	return &apiKeyActivityRecords, nil
}
