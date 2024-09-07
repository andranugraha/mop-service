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

func (c *cache) getMerchantKeyPatternByCode(code string) string {
	return fmt.Sprintf("merchant_data:merchantid:*:code:%s", code)
}

func (c *cache) getMerchantKeyPatternByID(id uint64) string {
	return fmt.Sprintf("merchant_data:merchantid:%d:*", id)
}

func (c *cache) getMerchantKey(id uint64, code string) string {
	return fmt.Sprintf("merchant_data:merchantid:%d:code:%s", id, code)
}

func (c *cache) GetMerchantByID(ctx context.Context, id uint64) (*Merchant, error) {
	merchant := &Merchant{}
	keyPattern := c.getMerchantKeyPatternByID(id)
	keys, err := c.client.Keys(ctx, keyPattern).Result()
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", err)
		return nil, err
	}

	if len(keys) == 0 {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", errors.New("merchant not found"))
		return nil, nil
	}

	key := keys[0]
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &merchant)
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to unmarshal merchant: %v", err)
		return nil, err
	}

	return merchant, nil
}

func (c *cache) GetMerchantByCode(ctx context.Context, code string) (*Merchant, error) {
	merchant := &Merchant{}
	keyPattern := c.getMerchantKeyPatternByCode(code)
	keys, err := c.client.Keys(ctx, keyPattern).Result()
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", err)
		return nil, err
	}

	if len(keys) == 0 {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", errors.New("merchant not found"))
		return nil, nil
	}

	key := keys[0]
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to get merchant: %v", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &merchant)
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_cache", "failed to unmarshal merchant: %v", err)
		return nil, err
	}

	return merchant, nil
}

func (c *cache) SetMerchant(ctx context.Context, merchant *Merchant) error {
	key := c.getMerchantKey(merchant.GetId(), merchant.GetCode())
	val, err := json.Marshal(merchant)
	if err != nil {
		logger.Error(ctx, "set_merchant_to_cache", "failed to marshal merchant: %v", err)
		return err
	}

	expiryInSeconds := config.GetCacheMerchantExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, val, expiryDuration).Err()
	if err != nil {
		logger.Error(ctx, "set_merchant_to_cache", "failed to set merchant: %v", err)
		return err
	}

	return nil
}

func (c *cache) GetMerchantOverview(ctx context.Context, code string) (*Merchant, error) {
	KEY := "merchant_overview:" + code
	merchant := &Merchant{}

	val, err := c.client.Get(ctx, KEY).Result()
	if err == redis.Nil {
		logger.Info(
			ctx, "fetch_merchant_overview_from_cache", "merchant not found in cache, fetching from db",
		)
		return nil, nil
	}
	if err != nil {
		logger.Error(
			ctx, "fetch_merchant_overview_from_cache", "failed to get merchant overview: %v", err,
		)
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &merchant)
	if err != nil {
		logger.Error(
			ctx, "fetch_merchant_overview_from_cache", "failed to unmarshal merchant overview: %v", err,
		)
		return nil, err
	}

	return merchant, nil
}

func (c *cache) SetMerchantOverview(ctx context.Context, code string, merchant *Merchant) error {
	key := "merchant_overview:" + code

	val, err := json.Marshal(merchant)
	if err != nil {
		logger.Error(
			ctx, "set_merchant_overview_to_cache", "failed to marshal merchant overview: %v", err,
		)
		return err
	}

	expiryInSeconds := config.GetCacheMerchantExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, val, expiryDuration).Err()
	if err != nil {
		logger.Error(
			ctx, "set_merchant_overview_to_cache", "failed to set merchant overview: %v", err,
		)
		return err
	}

	return nil
}
