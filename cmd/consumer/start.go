package consumer

import (
	"context"
	"fmt"
	"time"

	"reprocessing-engine/internal/events"
	"reprocessing-engine/internal/worker"
)

// runConsumer starts the consumer application
func RunConsumer(ctx context.Context) {
	fmt.Println("[consumer] Starting the consumer process")

	eventsChan := make(chan string)
	go events.PublishEventsToNATS(ctx, eventsChan)

	// Here you would implement the consumer logic
	// For example, subscribing to a NATS subject

	// Simulating work
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[consumer] Context cancelled, stopping consumer")
			return
		case <-ticker.C:
			fmt.Println("[consumer] Processing new message batch")
			// Simulate processing with a worker for a specific message
			go func() {
				// This simulates processing orders from a received message
				worker.StartWorkerPool(ctx, 1, 1, eventsChan)
			}()
		}
	}
}
