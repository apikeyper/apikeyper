package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var (
	dbInstance *service
	dbConfig   *DbConfig
)

type DbConfig struct {
	dbHost string
	dbUrl  string
	dbUser string
	dbPass string
	dbName string
	dbPort string
}

func SetupDb() *gorm.DB {

	dbConfig := GetDbConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		dbConfig.dbHost,
		dbConfig.dbUser,
		dbConfig.dbPass,
		dbConfig.dbName,
		dbConfig.dbPort,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(
		// &User{},
		&Workspace{},
		&RootKey{},
		&Api{},
		&ApiKey{},
	)

	return db
}

func GetDbConfig() *DbConfig {

	if dbConfig != nil {
		return dbConfig
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file to load")
	}

	var (
		// appEnv     = os.Getenv("APP_ENV")
		dbUrl  = os.Getenv("DATABASE_URL")
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
		dbPort = os.Getenv("DB_PORT")
	)

	return &DbConfig{
		dbHost: dbHost,
		dbUrl:  dbUrl,
		dbUser: dbUser,
		dbPass: dbPass,
		dbName: dbName,
		dbPort: dbPort,
	}
}
