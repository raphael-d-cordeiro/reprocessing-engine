package scheduler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"reprocessing-engine/internal/events"
	"reprocessing-engine/internal/worker"
)

const (
	maxWorkers       = 10
	scheduleInterval = 60 * time.Minute
)

// FetchTotalPages simulates a request to an API and returns the total number of available pages for reprocessing.
func fetchTotalPages(ctx context.Context) (int, error) {
	// Simulate HTTP request with context-based timeout
	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.exemplo.com/orders?page=1", nil)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to create request: %w", err)
	// }

	// Here you would perform the real API call. Simulating delay:
	time.Sleep(300 * time.Millisecond)

	// Handle context cancellation
	select {
	case <-ctx.Done():
		return 0, errors.New("request canceled by context")
	default:
	}

	// Simulate API response with total page count
	totalPages := 5
	fmt.Println("[scanner] Total pages from API:", totalPages)
	return totalPages, nil
}

// runScanner starts the scanner application
func RunScanner(ctx context.Context) {
	fmt.Println("[scanner] Starting Scanner application")

	eventsChan := make(chan string)
	go events.PublishEventsToNATS(ctx, eventsChan)

	ticker := time.NewTicker(scheduleInterval)
	defer ticker.Stop()

	runWorker := func() {
		totalPages, err := fetchTotalPages(ctx)
		if err != nil {
			fmt.Println("[scanner] Error fetching pages:", err)
			return
		}
		fmt.Println("[scanner] Starting worker pool")
		worker.StartWorkerPool(ctx, totalPages, maxWorkers, eventsChan)
	}

	runWorker() // Run immediately on startup

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[scanner] Gracefully shutting down...")
			return
		case <-ticker.C:
			runWorker()
		}
	}
}
