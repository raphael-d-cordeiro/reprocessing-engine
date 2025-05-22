package scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// FetchTotalPages simulates a request to an API and returns the total number of available pages for reprocessing.
func FetchTotalPages(ctx context.Context) (int, error) {
	// Simulate HTTP request with context-based timeout
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.exemplo.com/orders?page=1", nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

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