package events

import (
	"context"
	"encoding/json"
	"fmt"
	"keyify/internal/database"
	"log/slog"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func Consumer(ctx context.Context, config *QueueConfig, consumerName string) {

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "",
		DB:       0,
	})

	dbService := database.New()

	slog.Info(fmt.Sprintf("Starting consumer: %s", consumerName))

	for {
		result, err := client.BRPop(ctx, 0, config.QueueName).Result()
		if err != nil {
			fmt.Printf("[%s] error popping from queue: %v\n", consumerName, err)
			continue
		}

		fmt.Printf("[%s] received: %v\n", consumerName, result)

		// Convert the result to json
		eventPayload := EventPayload{}

		// Unmarshal the event payload
		err = json.Unmarshal([]byte(result[1]), &eventPayload)

		if err != nil {
			slog.Error(fmt.Sprintf("Error unmarshalling event payload: %v", err))
			continue
		}

		var usage string

		switch eventPayload.EventType {
		case API_KEY_VERIFY_SUCCESS:
			usage = "success"
		case API_KEY_VERIFY_FAILED:
			usage = "failed"
		case API_KEY_RATE_LIMITED:
			usage = "rate_limited"
		default:
			usage = "unknown"
		}

		// Log the API key usage
		apiKeyId := eventPayload.Data.ApiKeyId
		apiKeyUsage := database.ApiKeyUsage{
			ApiKeyId: uuid.MustParse(apiKeyId),
			Usage:    usage,
		}
		dbService.LogApiKeyUsage(&apiKeyUsage)
		slog.Info(fmt.Sprintf("Processed event: %v", eventPayload.EventType))
	}

}
