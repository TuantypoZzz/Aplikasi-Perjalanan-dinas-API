package configuration

import (
	"context"
	"golang-todo-app/exception"
	"os"

	"github.com/go-redis/redis/v8"
)

// InitRedis menginisialisasi koneksi Redis
func NewRedis() *redis.Client {
	addr := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Cek koneksi Redis
	_, err := RedisClient.Ping(context.Background()).Result()
	exception.PanicLogging(err)

	return RedisClient
}
