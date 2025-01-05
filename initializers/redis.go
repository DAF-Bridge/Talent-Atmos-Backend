package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectToRedis() *redis.Client {
	// Connect to Redis
	rs := os.Getenv("REDIS_URL")

	opt, _ := redis.ParseURL(rs)
	client := redis.NewClient(opt)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to ping to Redis: %v!\n", err)
	}

	fmt.Println("Successfully connected to Redis!")

	return client
}
