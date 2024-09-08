package server

import (
	"database/sql"
	"time"

	"apikeyper/internal/database"

	"github.com/google/uuid"
)

type CreateWorkspaceResponse struct {
	WorkspaceId uuid.UUID `json:"workspaceId"`
}

type FetchWorkspaceResponse struct {
	WorkspaceId   uuid.UUID `json:"workspaceId"`
	WorkspaceName string    `json:"workspaceName"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type CreateRootKeyResponse struct {
	RootKey string `json:"rootKey"`
}

type FetchRootKeyResponse struct {
	RootKeyId   uuid.UUID `json:"rootKeyId"`
	RootKeyName *string   `json:"rootKeyName,omitempty"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateApiResponse struct {
	ApiId uuid.UUID `json:"apiId"`
}

type CreateKeyResponse struct {
	ApiId uuid.UUID `json:"apiId"`
	KeyId uuid.UUID `json:"keyId"`
	Key   string    `json:"key"`
}

type FetchApiKeyResponse struct {
	ApiKeyId  uuid.UUID    `json:"apiKeyId"`
	ApiId     uuid.UUID    `json:"apiId"`
	Name      *string      `json:"name"`
	ExpiresAt sql.NullTime `json:"expiresAt,omitempty"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
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
