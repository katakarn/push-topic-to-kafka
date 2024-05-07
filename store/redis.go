package store

import (
	"context"
	"log"
	"testKafka/config"
	"strings"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) *redis.ClusterClient {

	addrs := strings.Split(cfg.RedisClusterServer, ",")

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		WriteTimeout: -1,
		ReadTimeout:  -1,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Print("Error connecting to Redis:", err)
		panic(err)
	}

	log.Print("Connected to Redis")
	return client
}
