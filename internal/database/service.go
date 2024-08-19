package database

import (
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	Health() map[string]string

	// Workspace
	CreateWorkspace(workspace *Workspace) (uuid.UUID, error)

	// RootKey
	CreateRootKey(rootKey *RootKey) (uuid.UUID, error)
	FetchRootKey(rootHashedKey string) (*RootKey, error)

	// Api
	CreateApi(api *Api) (uuid.UUID, error)
	FetchApi(apiId uuid.UUID) (*Api, error)
	ListApis(workspaceId uuid.UUID) (*[]Api, error)

	// ApiKey
	CreateApiKey(apiKey *ApiKey) (uuid.UUID, error)
	FetchApiKeyById(apiKeyId string) (*ApiKey, error)
	VerifyApiKey(apiKeyHashed string) (*ApiKey, error)
	UpdateApiKeyStatus(apiKeyId uuid.UUID, status string) (*ApiKey, error)

	// ApiKeyUsage
	LogApiKeyUsage(apiKeyUsage *ApiKeyUsage) (uuid.UUID, error)
}

type service struct {
	db *gorm.DB
}

func New() Service {
	// Reuse Connection
	if dbService != nil {
		return dbService
	}

	db := SetupDb()

	dbService = &service{
		db: db,
	}
	return dbService
}
