package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func (c *cache) getUserKeyPatternByEmail(email string) string {
	return fmt.Sprintf("user_data:userid:*:email:%s", email)
}

func (c *cache) getUserKey(userId uint64, email string) string {
	return fmt.Sprintf("user_data:userid:%d:email:%s", userId, email)
}

func (c *cache) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	keyPattern := c.getUserKeyPatternByEmail(email)
	keys, err := c.client.Keys(ctx, keyPattern).Result()
	if err != nil {
		logger.Error(ctx, "fetch_user_from_cache", "failed to fetch user: %v", err)
		return nil, err
	}

	if len(keys) == 0 {
		logger.Warn(ctx, "fetch_user_from_cache", "failed to fetch user: user not found in redis")
		return nil, nil
	}

	jsonValue, err := c.client.Get(ctx, keys[0]).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Warn(ctx, "fetch_user_from_cache", "failed to fetch user: user not found in redis")
			return nil, nil
		}
		logger.Error(ctx, "fetch_user_from_cache", "failed to fetch user: %v", err)
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal([]byte(jsonValue), user)
	if err != nil {
		logger.Error(ctx, "fetch_user_from_cache", "failed to unmarshal user: %v", err)
		return nil, err
	}

	return user, nil
}

func (c *cache) SetUser(ctx context.Context, user *User) error {
	key := c.getUserKey(user.GetId(), user.GetEmail())
	jsonValue, err := json.Marshal(user)
	if err != nil {
		logger.Error(ctx, "set_user_to_cache", "failed to marshal user: %v", err)
		return err
	}

	expiryInSeconds := config.GetCacheUserExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, jsonValue, expiryDuration).Err()
	if err != nil {
		logger.Error(ctx, "set_user_to_cache", "failed to set user: %v", err)
		return err
	}

	logger.Data(ctx, "set_user_to_cache", "user: %+v", user)
	return nil
}
