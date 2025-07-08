package config

import (
	"context"
	"fmt"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	client *redis.Client
}

func ConnectRedis(config RedisConfig) *Cache {
	client := redis.NewClient(getRedisConfig(config))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Fatal("Failed to connect to Redis:", err)
	}

	return &Cache{
		client: client,
	}
}

func getRedisConfig(config RedisConfig) *redis.Options {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	return &redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	}
}

func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}
