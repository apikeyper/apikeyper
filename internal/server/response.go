package server

import (
	"time"

	"apikeyper/internal/database"

	"github.com/google/uuid"
)

type CreateWorkspaceResponse struct {
	WorkspaceId uuid.UUID `json:"workspaceId"`
}

type CreateRootKeyResponse struct {
	RootKey string `json:"rootKey"`
}

type CreateApiResponse struct {
	ApiId uuid.UUID `json:"apiId"`
}

type CreateKeyResponse struct {
	ApiId uuid.UUID `json:"apiId"`
	KeyId uuid.UUID `json:"keyId"`
	Key   string    `json:"key"`
}

type FetchApiResponse struct {
	ApiId       uuid.UUID `json:"apiId"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	ApiName     string    `json:"apiName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type VerifyApiKeyResponse struct {
	KeyId uuid.UUID `json:"keyId"`
	ApiId uuid.UUID `json:"apiId"`
	Valid bool      `json:"valid"`
}

type RevokeApiKeyResponse struct {
	KeyId   uuid.UUID `json:"keyId"`
	ApiId   uuid.UUID `json:"apiId"`
	Revoked bool      `json:"revoked"`
}

type FetchApiKeyUsageResponse struct {
	ApiKeyId uuid.UUID                   `json:"apiKeyId"`
	Records  []database.ApiKeyUsageCount `json:"records"`
	Count    int                         `json:"count"`
}
