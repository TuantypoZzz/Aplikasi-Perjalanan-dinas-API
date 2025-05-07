package repository

import "time"

type CacheRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, entry interface{}, expiration time.Duration) error
	Delete(key string) error
}
