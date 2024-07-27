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

var (
	appEnv           = os.Getenv("APP_ENV")
	tursoDbUrl       = os.Getenv("TURSO_DATABASE_URL")
	tursoDbAuthToken = os.Getenv("TURSO_AUTH_TOKEN")
	dbInstance       *service
)

func SetupDb(dbDriver string, dbUri string) *gorm.DB {
	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: dbDriver,
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

func GetDbConfig() (string, string) {
	var dbDriver string
	var dbUri string
	if appEnv == "test" {
		dbDriver = "sqlite3"
		dbUri = "file::memory:?cache=shared"
	} else {
		dbDriver = "libsql"
		dbUri = fmt.Sprintf("%s?authToken=%s", tursoDbUrl, tursoDbAuthToken)
	}

	return dbDriver, dbUri
}
