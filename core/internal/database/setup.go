package database

import (
	"database/sql"
	"log"
	"log/slog"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var (
	dbInstance *gorm.DB
	dbService  *service
	dbConfig   *DbConfig
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

	sqlDB, _ := sql.Open("pgx", dsn)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	var dbConnErr error
	dbInstance, dbConnErr = gorm.Open(
		postgres.New(postgres.Config{
			Conn: sqlDB,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})

	if dbConnErr != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(dbConnErr)
	}

	return dbInstance
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
