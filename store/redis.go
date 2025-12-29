package store

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	RDB *redis.Client
)

func InitRedis() error {
	// RDB = redis.NewClient(&redis.Options{
	// 	Addr:         os.Getenv("REDIS_ADDR"),
	// 	Password:     os.Getenv("REDIS_PASSWORD"),
	// 	DB:           0,
	// 	ReadTimeout:  2 * time.Second,
	// 	WriteTimeout: 2 * time.Second,
	// })

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	client := redis.NewClient(opt)
	RDB = client
	return RDB.Ping(Ctx).Err()
}
