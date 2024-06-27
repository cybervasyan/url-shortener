package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-11145.c261.us-east-1-4.ec2.redns.redis-cloud.com:11145",
		Username: "default",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error occured while Redis initialization: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient

	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed while saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveOriginalUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}

	return result
}
