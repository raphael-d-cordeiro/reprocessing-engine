package cache

import (
	"context"
	"crypto/tls"
	"fmt"

	redislib "github.com/redis/go-redis/v9"
)

type CacheClient struct {
	Client *redislib.Client
}

func New(host string, port int, password string) (*CacheClient, error) {
	opts := redislib.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		DB:       0,
		PoolSize: 300,
	}
	if password != "" {
		opts.Password = password // set password if provided
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: true, // For development purposes only
		}
	}
	client := redislib.NewClient(&opts)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	return &CacheClient{
		Client: client,
	}, nil
}

func (c *CacheClient) Get(ctx context.Context, key string) (string, error) {
	value, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redislib.Nil {
			return "", nil // Key does not exist
		}
		return "", fmt.Errorf("failed to get cache key %s: %w", key, err)
	}
	return value, nil
}
func (c *CacheClient) Set(ctx context.Context, key string, value interface{}, expiration int64) error {
	_, err := c.Client.Set(ctx, key, value, 0).Result()
	if err != nil {
		return fmt.Errorf("failed to set cache key %s: %w", key, err)
	}
	return nil
}

func (c *CacheClient) SetNx(ctx context.Context, key string, value interface{}, expiration int64) (bool, error) {
	result, err := c.Client.SetNX(ctx, key, value, 0).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set cache key %s: %w", key, err)
	}
	return result, nil
}

func (c *CacheClient) Delete(ctx context.Context, key string) error {
	_, err := c.Client.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to delete cache key %s: %w", key, err)
	}
	return nil
}
