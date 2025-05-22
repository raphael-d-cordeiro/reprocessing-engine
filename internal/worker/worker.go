package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"reprocessing-engine/internal/models"
)

// StartWorkerPool limits the number of concurrent workers processing pages.
func StartWorkerPool(ctx context.Context, totalPages, maxWorkers int, eventChan chan<- string) {
	pageJobs := make(chan int)
	var wg sync.WaitGroup

	// Start fixed number of workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for page := range pageJobs {
				processPage(ctx, page, eventChan, workerID)
			}
		}(i + 1)
	}

	// Send page jobs
	for page := 1; page <= totalPages; page++ {
		select {
		case <-ctx.Done():
			fmt.Println("[worker-pool] Context canceled before dispatching all pages")
			close(pageJobs)
			wg.Wait()
			return
		case pageJobs <- page:
		}
	}
	close(pageJobs)
	wg.Wait()
}

// processPage simulates fetching and processing orders from a specific page.
func processPage(ctx context.Context, page int, eventChan chan<- string, workerID int) {
	fmt.Printf("[worker-%d] Processing page %d\n", workerID, page)

	orders := []models.Order{
		{ID: fmt.Sprintf("order-%d-a", page)},
		{ID: fmt.Sprintf("order-%d-b", page)},
	}

	for _, order := range orders {
		time.Sleep(100 * time.Millisecond)

		select {
		case <-ctx.Done():
			fmt.Printf("[worker-%d] Context canceled while processing page %d\n", workerID, page)
			return
		case eventChan <- fmt.Sprintf("proposal.processing.failed: %s", order.ID):
			fmt.Printf("[worker-%d] Event emitted for order %s\n", workerID, order.ID)
		}
	}
}
