package config

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var globalCache *redis.Client

func initCache() error {
	globalCache = redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Username:     redisUsername,
		Password:     redisPassword,
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 20,
		IdleTimeout:  5 * time.Minute,
		MaxConnAge:   30 * time.Minute,
		PoolTimeout:  10 * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	_, err := globalCache.Ping(context.Background()).Result()
	if err != nil {
		return errors.New("failed to connect to cache redis: " + err.Error())
	}

	return nil
}

func GetCache() *redis.Client {
	return globalCache
}
