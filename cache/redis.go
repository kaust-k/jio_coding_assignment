package cache

import (
	"log"

	"github.com/go-redis/redis/v7"

	"jwt_server/config"
)

var redisClient *redis.Client

func init() {
	cfg := config.GetConfig()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Printf("REDIS Ping response:: %s\n", pong)
}

// GetRedisClient gets redis client
func GetRedisClient() *redis.Client {
	return redisClient
}
