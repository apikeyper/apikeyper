package database

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Create api key
	CreateApiKey(apiKey *ApiKey) string
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
