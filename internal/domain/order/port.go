package order

import "context"

type OrderRepository interface {
	// FetchOrders searches for orders based on the given page number.
	RetrieveOrders(ctx context.Context, page int, partnerKey string) ([]Order, error)
}

type OrderPublisher interface {
	// Publish sends an event related to order processing.
	Publish(ctx context.Context, order Order) error
}

// OrderDTO representa o Data Transfer Object para Order (usado para transporte/serialização).
type Order struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	ProductKey string `json:"product_key"`
}

type OrdersResponse struct {
	Data       []Order `json:"data"`
	TotalPages int     `json:"total_pages"`
}
