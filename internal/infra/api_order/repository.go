package api_order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reprocessing-engine/internal/domain/order"
	"time"
)

type APIOrderRepository struct {
	baseURL string
	client  *http.Client
}

func NewAPIOrderRepository(baseURL string) *APIOrderRepository {
	return &APIOrderRepository{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (r *APIOrderRepository) FetchOrders(ctx context.Context, page int) ([]order.Order, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", r.baseURL+"/orders", nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	var response OrdersResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("Failed to decode response: %w", err)
	}
	orders := make([]order.Order, 0, len(response.Data))
	for _, dto := range response.Data {
		orders = append(orders, order.Order{
			ID:         dto.ID,
			Status:     dto.Status,
			ProductKey: dto.ProductKey,
		})
	}
}
