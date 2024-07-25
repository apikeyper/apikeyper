package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
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

var (
	tursoDbUrl       = os.Getenv("TURSO_DATABASE_URL")
	tursoDbAuthToken = os.Getenv("TURSO_AUTH_TOKEN")
	dbInstance       *service
)

func SetupDb(dbUri string) *gorm.DB {
	// Setup database
	log.Fatal(dbUri)
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "libsql",
		DSN:        dbUri,
	}), &gorm.Config{})

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&ApiKey{})

	return db
}

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	var env = os.Getenv("APP_ENV")

	var dbUri string
	if env == "local" {
		dbUri = "file::memory:?cache=shared"
	} else {
		dbUri = fmt.Sprintf("%s?authToken=%s", tursoDbUrl, tursoDbAuthToken)
	}

	log.Default().Println("Connecting to database:", dbUri)

	db := SetupDb(dbUri)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}
