package main

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Set(key string, val interface{}) error
	Get(key string, val interface{}) (bool, error)
}

type RedisCache struct {
	client     *redis.Client
	expiration time.Duration
}

func NewRedisCache(opts *redis.Options, expiration time.Duration) *RedisCache {
	return &RedisCache{
		client:     redis.NewClient(opts),
		expiration: expiration,
	}
}

func (cache *RedisCache) Set(key string, val interface{}) error {
	json, err := json.Marshal(val)

	if err != nil {
		return err
	}

	cache.client.Set(key, string(json), cache.expiration*time.Second)
	return nil
}

func (cache *RedisCache) Get(key string, val interface{}) (bool, error) {
	data, err := cache.client.Get(key).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(data), val)

	if err != nil {
		return false, err
	}

	return true, nil
}
