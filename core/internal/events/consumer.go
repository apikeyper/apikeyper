package events

import (
	"apikeyper/internal/common"
	"apikeyper/internal/database"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func Consumer(ctx context.Context, config *common.RedisConfig, consumerName string) {

	client := common.NewRedisClient(config, "events")

	dbService := database.New()

	slog.Info(fmt.Sprintf("Starting consumer: %s", consumerName))

	for {
		result, err := client.BRPopLPush(ctx, eventQueueName, eventTempQueueName, 0).Result()

		if err != nil {
			fmt.Printf("[%s] error popping from queue: %v\n", consumerName, err)
			continue
		}
		slog.Info(fmt.Sprintf("[%s] received: %v\n", consumerName, result))

		// Convert the result to json
		eventPayload := EventPayload{}

		// Unmarshal the event payload
		err = json.Unmarshal([]byte(result), &eventPayload)

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
		case API_KEY_REVOKED:
			usage = "revoked"
		default:
			slog.Info(fmt.Sprintf("Skipping unknown event type: %v", eventPayload.EventType))
			client.LRem(ctx, "tempQ", 0, result)
			continue
		}

		// Log the API key usage
		apiKeyId := eventPayload.Data.ApiKeyId
		apiKeyUsage := database.ApiKeyActivity{
			ID:       eventPayload.Data.EventId,
			ApiKeyId: uuid.MustParse(apiKeyId),
			Usage:    usage,
		}
		dbService.LogApiKeyUsage(&apiKeyUsage)

		client.LRem(ctx, "tempQ", 0, result)

		slog.Info(fmt.Sprintf("Processed event: %v", eventPayload.EventType))
	}

}
