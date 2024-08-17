package main

import (
	"context"
	"keyify/internal/events"
	"log/slog"
)

func main() {
	ctx := context.Background()

	// Start the worker and keep it running
	slog.Info("Starting worker")

	config := events.GetQueueConfig()

	go events.Consumer(ctx, config, "1")
	go events.Consumer(ctx, config, "2")
	go events.Consumer(ctx, config, "3")

	select {}
}
