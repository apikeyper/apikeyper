package events

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

func (s *messageService) Publish(ctx context.Context, eventPayload EventPayload) error {
	// Push the event to the queue
	parsedPayload, marshalErr := json.Marshal(eventPayload)

	if marshalErr != nil {
		slog.Error(fmt.Sprintf("error marshalling event payload: %v", marshalErr))
		return fmt.Errorf("error marshalling event payload: %v", marshalErr)
	}

	_, pushResult := s.client.LPush(ctx, s.queueName, parsedPayload).Result()
	if pushResult != nil {
		slog.Error(fmt.Sprintf("error pushing to queue: %v", pushResult))
		return fmt.Errorf("error pushing to queue: %v", pushResult)
	}

	slog.Info(
		fmt.Sprintf(
			"event: %s for workspace %s pushed to queue",
			eventPayload.EventType,
			eventPayload.Data.WorkspaceId,
		),
	)

	return nil
}
