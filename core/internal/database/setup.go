package database

import (
	"log"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var (
	dbService *service
	dbConfig  *DbConfig
)

type DbConfig struct {
	dbUrl          string
	dbDebugLogging bool
}

func ParseDbUrl(dbConfig *DbConfig) string {
	if dbConfig.dbUrl != "" {
		return dbConfig.dbUrl
	}

	panic("DATABASE_URL is not set")
}

func GetGormDb() *gorm.DB {
	dbConfig := GetDbConfig()
	dsn := ParseDbUrl(dbConfig)

	var logLevel logger.LogLevel
	if dbConfig.dbDebugLogging {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	return db
}

func SetupDb() *gorm.DB {

	db := GetGormDb()

	// Migrate the schema
	db.AutoMigrate(
		&User{},
		&Session{},
		&Workspace{},
		&RootKey{},
		&Api{},
		&ApiKey{},
		&ApiKeyActivity{},
		&ApiKeyRateLimitConfig{},
	)

	return db
}

func GetDbConfig() *DbConfig {

	if dbConfig != nil {
		return dbConfig
	}

	err := godotenv.Load()
	if err != nil {
		slog.Debug("No .env file to load")
	}

	var (
		dbUrl          = os.Getenv("DATABASE_URL")
		dbDebugLogging = os.Getenv("DATABASE_DEBUG_LOGGING") == "true"
	)

	return &DbConfig{
		dbUrl:          dbUrl,
		dbDebugLogging: dbDebugLogging,
	}
}
