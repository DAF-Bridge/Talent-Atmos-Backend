package service

import (
	"log"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/ports"
)

type RedisService struct {
	RedisClient ports.RedisClient
}

func NewRedisService(redisClient ports.RedisClient) *RedisService {
	return &RedisService{
		RedisClient: redisClient,
	}
}

func (rs *RedisService) CacheData(key string, value string) error {
	err := rs.RedisClient.Set(key, value)

	if err != nil {
		log.Printf("Error setting key %s in redis: %v", key, err)
		return err
	}

	return nil
}

func (rs *RedisService) FetchData(key string) (string, error) {
	value, err := rs.RedisClient.Get(key)
	
	if err != nil {
		log.Printf("Error fetching key %s from redis: %v", key, err)
		return "", err
	}

	return value, nil
}
