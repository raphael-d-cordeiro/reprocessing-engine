// events.go handles processing and publishing events
package events

import (
	"context"
	"fmt"
	"time"
)

// PublishEventsToNATS consumes events from the channel and publishes them to NATS
func PublishEventsToNATS(ctx context.Context, eventsChan <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[events] Context cancelled, stopping event publisher")
			return
		case event, ok := <-eventsChan:
			if !ok {
				fmt.Println("[events] Channel closed, stopping event publisher")
				return
			}
			// Simulate publishing to NATS
			fmt.Printf("[events] Publishing event to NATS: %s\n", event)
			time.Sleep(50 * time.Millisecond)
		}
	}
}
