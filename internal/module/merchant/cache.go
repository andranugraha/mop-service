package merchant

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

func (c *cache) getMerchantKey(code string) string {
	return fmt.Sprintf("merchant_data:code:%s", code)
}

func (c *cache) GetMerchantByCode(ctx context.Context, code string) (*Merchant, error) {
	key := c.getMerchantKey(code)
	jsonValue, err := c.client.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Warn(ctx, "fetch_merchant_from_cache", "failed to fetch merchant: merchant not found in redis")
			return nil, nil
		}

		logger.Error(ctx, "fetch_merchant_from_cache", "failed to fetch merchant: %v", err)
		return nil, err
	}

	merchant := &Merchant{}
	err = json.Unmarshal([]byte(jsonValue), merchant)

	if err != nil {
		logger.Warn(ctx, "fetch_merchant_from_cache", "failed to unmarshal merchant: %v", err)
		return nil, err
	}

	logger.Data(ctx, "fetch_merchant_from_cache", "merchant fetched from cache: %v", merchant)
	return merchant, nil
}

func (c *cache) SetMerchant(ctx context.Context, merchant *Merchant) error {
	key := c.getMerchantKey(merchant.GetCode())
	jsonValue, err := json.Marshal(merchant)

	if err != nil {
		logger.Error(ctx, "set_merchant_to_cache", "failed to marshal merchant: %v", err)
		return err
	}

	expiryInSeconds := config.GetCacheUserExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, jsonValue, expiryDuration).Err()

	if err != nil {
		logger.Error(ctx, "set_merchant_to_cache", "failed to set merchant: %v", err)
		return err
	}

	logger.Data(ctx, "set_merchant_to_cache", "merchant set to cache: %v", merchant)
	return nil
}
