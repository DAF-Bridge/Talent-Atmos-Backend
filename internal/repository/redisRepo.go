package repository

import "github.com/go-redis/redis/v8"

type RedisRepository struct {
	client *redis.Client
}

