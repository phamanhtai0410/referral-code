package cache

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/pkg/utils"
	"example.com/refcode/v1/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	option := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}
	client = redis.NewClient(option)
	InitCodeCounter()
	log.Println("[REDIS] connection successful")
}

func InitCodeCounter() {
	_, err := client.Get(context.TODO(), constants.CacheCounter).Int64()
	if err == redis.Nil {
		var latest models.User
		collection := database.GetCollection(models.TableDetails)
		options := options.FindOne().SetSort(bson.M{"id": -1})
		err = collection.FindOne(context.TODO(), bson.M{}, options).Decode(&latest)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				client.Set(context.TODO(), constants.CacheCounter, 111111, 0)
			} else {
				log.Fatal("[REDIS #1]", err)
			}
		} else {
			client.Set(context.TODO(), constants.CacheCounter, latest.Id, 0)
		}

	}
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
	InitCodeCounter()
	ctx, _ := database.NewContext()
	return client.Incr(ctx, key).Result()
}
