package banner

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func (c *cache) getBannersKeyPatternByMerchantID(merchantID uint64) string {
	return fmt.Sprintf("banners:merchantid:%d", merchantID)
}

func (c *cache) GetActiveBannersByMerchantID(ctx context.Context, merchantID uint64) ([]*Banner, error) {
	key := c.getBannersKeyPatternByMerchantID(merchantID)
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "get_active_banners_by_merchant_id", "failed to get banners: %v", err.Error())
		return nil, err
	}

	var banners []*Banner
	err = json.Unmarshal([]byte(val), &banners)
	if err != nil {
		logger.Error(ctx, "get_active_banners_by_merchant_id", "failed to unmarshal banners: %v", err.Error())
		return nil, err
	}

	logger.InfoWithData(ctx, "get_active_banners_by_merchant_id", "banners: %v", banners)
	return banners, nil
}

func (c *cache) SetActiveBannersByMerchantID(ctx context.Context, merchantID uint64, banners []*Banner) error {
	key := c.getBannersKeyPatternByMerchantID(merchantID)
	val, err := json.Marshal(banners)
	if err != nil {
		logger.Error(ctx, "set_active_banners_by_merchant_id", "failed to marshal banners: %v", err.Error())
		return err
	}

	expiryInSeconds := config.GetCacheLandingBannersExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, val, expiryDuration).Err()
	if err != nil {
		logger.Error(ctx, "set_active_banners_by_merchant_id", "failed to set banners: %v", err.Error())
		return err
	}

	return nil
}
