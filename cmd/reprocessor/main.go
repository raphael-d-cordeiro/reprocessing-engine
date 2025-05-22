package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"reprocessing-engine/internal/events"
	"reprocessing-engine/internal/scanner"
	"reprocessing-engine/internal/worker"
)

const (
	maxWorkers       = 10
	scheduleInterval = 60 * time.Minute
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	eventsChan := make(chan string)
	go events.PublishEventsToNATS(ctx, eventsChan)

	ticker := time.NewTicker(scheduleInterval)
	defer ticker.Stop()

	runOnce := func() {
		totalPages, err := scanner.FetchTotalPages(ctx)
		if err != nil {
			fmt.Println("Error fetching pages:", err)
			return
		}
		fmt.Println("[main] Starting worker pool")
		worker.StartWorkerPool(ctx, totalPages, maxWorkers, eventsChan)
	}

	runOnce() // Run immediately on startup

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[main] Gracefully shutting down...")
			return
		case <-ticker.C:
			runOnce()
		}
	}
}
