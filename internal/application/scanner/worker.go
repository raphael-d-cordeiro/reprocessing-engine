package scanner

import (
	"context"

	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/domain/engine"
	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/domain/order"
)

type Scanner struct {
	cacheClient     engine.CacheInterface
	orderRepository order.OrderRepository
}

func New(cacheClient engine.CacheInterface, orderRepository order.OrderRepository) *Scanner {
	return &Scanner{
		cacheClient:     cacheClient,
		orderRepository: orderRepository,
	}
}

func (s *Scanner) Start(ctx context.Context) {
	// This is where the scanner logic would be implemented.
	// For example, it could periodically scan for new orders,
	// process them, and store results in the cache.
	// The actual implementation would depend on the specific requirements of the application.
}
