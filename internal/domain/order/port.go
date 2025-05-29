package order

import "context"

type OrderRepository interface {
	// FetchOrders searches for orders based on the given page number.
	FetchOrders(ctx context.Context, page int) ([]Order, error)
}

type OrderPublisher interface {
	// Publish sends an event related to order processing.
	Publish(ctx context.Context, order Order) error
}
