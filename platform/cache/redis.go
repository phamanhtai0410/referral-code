package cache

import (
	"log"
	"os"
	"strconv"
	"time"

	"example.com/refcode/v1/pkg/utils"
	"example.com/refcode/v1/platform/database"

	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.

var client *redis.Client

func init() {
	// Define Redis database number.
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		log.Fatal("[redis] Error creating connection")
	}

	// Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}
	client = redis.NewClient(options)
	log.Println("[REDIS] connection successful")
}

func Shutdown() {
	if client != nil {
		// ctx, _ := database.NewContext()
		client.Close()
		log.Println("[CACHE] Shutdown client connection")
	}
}

func Set(key string, data string, expire time.Duration) error {
	ctx, _ := database.NewContext()
	return client.Set(ctx, key, data, expire).Err()
}

func Get(key string) (int64, error) {
	ctx, _ := database.NewContext()
	return client.Get(ctx, key).Int64()
}

func Save(key string, data string) error {
	ctx, _ := database.NewContext()
	return client.Set(ctx, key, data, 0).Err()
}

func IsUnique(code string) bool {
	ctx, _ := database.NewContext()
	exists, err := client.Exists(ctx, code).Result()
	if err != nil {
		panic(err)
	}
	return exists == 0
}

func Incr(key string) (int64, error) {
	ctx, _ := database.NewContext()
	return client.Incr(ctx, key).Result()
}
