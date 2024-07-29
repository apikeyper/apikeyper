package database

import (
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	Health() map[string]string

	// RootKey
	CreateRootKey(rootKey *RootKey) uuid.UUID
	FetchRootKey(rootHashedKey string) *RootKey

	// Api
	CreateApi(api *Api) uuid.UUID
	FetchApi(apiId uuid.UUID) (*Api, error)

	// ApiKey
	CreateApiKey(apiKey *ApiKey) uuid.UUID
}

type service struct {
	db *gorm.DB
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	dbDriver, dbUri := GetDbConfig()

	db := SetupDb(dbDriver, dbUri)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}
