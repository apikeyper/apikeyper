package common

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisUrl string
}

func GetRedisConfig() *RedisConfig {
	err := godotenv.Load()
	if err != nil {
		slog.Debug("No .env file to load")
	}

	return &RedisConfig{
		RedisUrl: os.Getenv("REDIS_URL"),
	}
}

func NewRedisClient(config *RedisConfig, service string) *redis.Client {

	var db int
	switch service {
	case "events":
		db = 0
	case "ratelimit":
		db = 1
	default:
		slog.Info(fmt.Sprintf("Skipping unknown service: %v", service))
		return nil
	}

	return redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "",
		DB:       db,
	})
}
