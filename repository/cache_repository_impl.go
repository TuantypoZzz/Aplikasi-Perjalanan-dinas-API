package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type cacheRepositoryImpl struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) CacheRepository {
	return &cacheRepositoryImpl{client: client}
}

// Get implements CacheRepository.
func (r cacheRepositoryImpl) Get(key string) ([]byte, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("redis get error: %w", err)
	}
	return []byte(val), nil
}

func (r cacheRepositoryImpl) Set(key string, entry interface{}, expiration time.Duration) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshaling error: %w", err)
	}

	return r.client.Set(context.Background(), key, data, expiration).Err()
}

func (r cacheRepositoryImpl) Delete(key string) error {
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		return fmt.Errorf("redis delete error: %w", err)
	}
	return nil
}
