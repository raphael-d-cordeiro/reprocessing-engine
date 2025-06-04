package consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartConsumer(ctx context.Context) {
	// Initialize the consumer service
	// This is where you would set up your message queue consumer, e.g., RabbitMQ, Kafka, etc.
	log.Println("Starting consumer service...")

	// Simulate consumer work
	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer service shutting down gracefully...")
			return
		default:
			// Simulate processing messages
			log.Println("Processing messages...")
			time.Sleep(2 * time.Second) // Simulate work
		}
	}
}

func Run(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)

	go StartConsumer(ctx)
	HandleOSSignal(cancelFunc)
	log.Println("All services shutdown gracefully. Exiting... Consumer Service")

}

func HandleOSSignal(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	sig := <-signals
	signal.Stop(signals)
	log.Printf("Received signal: %s, Initiating graceful shutdown...\n", sig)
	cancel()
}
