package events

import (
	"context"

	"apikeyper/internal/common"

	"github.com/redis/go-redis/v9"
)

var messageServiceInstance *messageService

var (
	eventQueueName     = "queue"
	eventTempQueueName = "tempQ"
)

type MessageService interface {
	Publish(ctx context.Context, eventPayload EventPayload) error
}

type messageService struct {
	client    *redis.Client
	queueName string
}

func New() MessageService {

	if messageServiceInstance != nil {
		return messageServiceInstance
	}

	config := common.GetRedisConfig()

	messageServiceInstance = &messageService{
		client:    common.NewRedisClient(config, "events"),
		queueName: eventQueueName,
	}

	return messageServiceInstance
}
