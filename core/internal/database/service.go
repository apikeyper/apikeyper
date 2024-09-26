package database

import (
	"log/slog"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	Health() map[string]string

	// User
	CreateUser(user *User) (string, error)
	FetchUserByGithubId(githubId string) (*User, error)

	// Workspace
	CreateWorkspace(workspace *Workspace) (uuid.UUID, error)
	FetchWorkspaceById(workspaceId uuid.UUID) (*Workspace, error)

	// RootKey
	CreateRootKey(rootKey *RootKey) (uuid.UUID, error)
	FetchRootKey(rootHashedKey string) (*RootKey, error)
	FetchRootKeyById(rootKeyId uuid.UUID) (*RootKey, error)
	ListRootKeysForWorkspace(workspaceId uuid.UUID) (*[]RootKey, error)

	// Api
	CreateApi(api *Api) (uuid.UUID, error)
	FetchApi(apiId uuid.UUID) (*Api, error)
	ListApis(workspaceId uuid.UUID) (*[]Api, error)

	// ApiKey
	CreateApiKey(apiKey *ApiKey) (uuid.UUID, error)
	FetchApiKeyById(apiKeyId uuid.UUID) (*ApiKey, error)
	VerifyApiKey(apiKeyHashed string) (*ApiKey, error)
	UpdateApiKeyStatus(apiKeyId uuid.UUID, status string) (*ApiKey, error)
	ListApiKeysForApi(apiId uuid.UUID) (*[]ApiKey, error)

	// ApiKeyUsage
	LogApiKeyUsage(apiKeyUsage *ApiKeyActivity) (uuid.UUID, error)
	FetchApiKeyUsage(apiKeyId uuid.UUID, interval string) (*[]ApiKeyUsageCount, error)
}

type service struct {
	db *gorm.DB
}

func New() Service {
	// Reuse Connection
	if dbService != nil {
		slog.Info("Using existing db service")
		return dbService
	}

	if dbInstance == nil {
		slog.Info("Creating new db instance")
		dbInstance = SetupDb()
	} else {
		slog.Info("Using existing db instance")
	}

	dbService = &service{
		db: dbInstance,
	}
	return dbService
}
