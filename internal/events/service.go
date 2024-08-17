package events

import (
	"context"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var messageServiceInstance *messageService

type QueueConfig struct {
	RedisUrl  string
	QueueName string
}

func GetQueueConfig() *QueueConfig {
	err := godotenv.Load()
	if err != nil {
		slog.Debug("No .env file to load")
	}

	return &QueueConfig{
		RedisUrl:  os.Getenv("REDIS_URL"),
		QueueName: "queue",
	}
}

type MessageService interface {
	Publish(ctx context.Context, eventPayload EventPayload) error
}

func NewRedisClient(config *QueueConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "",
		DB:       0,
	})
}

type messageService struct {
	client    *redis.Client
	queueName string
}

func New() MessageService {

	if messageServiceInstance != nil {
		return messageServiceInstance
	}

	config := GetQueueConfig()

	messageServiceInstance = &messageService{
		client:    NewRedisClient(config),
		queueName: config.QueueName,
	}

	return messageServiceInstance
}
