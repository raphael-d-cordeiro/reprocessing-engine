package engine

import "context"

type CacheInterface interface {
	// Get retrieves an item from the cache by its key.
	Get(ctx context.Context, key string) (string, error)
	// Set stores an item in the cache with the specified key and value.
	Set(ctx context.Context, key string, value interface{}, expiration int64) error
	// SetNx sets an item in the cache only if the key does not already exist.
	SetNx(ctx context.Context, key string, value interface{}, expiration int64) (bool, error)
	// Delete removes an item from the cache by its key.
	Delete(ctx context.Context, key string) error
}
