package additionalfee

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func (c *cache) getMerchantAdditionalFeeKey(merchantID uint64) string {
	return fmt.Sprintf("additional_fee:merchant_id:%d", merchantID)
}

func (c *cache) GetActiveAdditionalFeesByMerchantID(ctx context.Context, merchantID uint64) ([]*AdditionalFee, error) {
	key := c.getMerchantAdditionalFeeKey(merchantID)
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_additional_fee_from_cache", "failed to get additional fee: %v", err)
		return nil, err
	}

	var additionalFee []*AdditionalFee
	err = json.Unmarshal([]byte(val), &additionalFee)
	if err != nil {
		logger.Error(ctx, "fetch_additional_fee_from_cache", "failed to unmarshal additional fee: %v", err)
		return nil, err
	}

	return additionalFee, nil
}

func (c *cache) SetAdditionalFeeByMerchantID(ctx context.Context, merchantID uint64, additionalFee []*AdditionalFee) error {
	key := c.getMerchantAdditionalFeeKey(merchantID)
	val, err := json.Marshal(additionalFee)
	if err != nil {
		logger.Error(ctx, "store_additional_fee_to_cache", "failed to marshal additional fee: %v", err)
		return err
	}

	if err = c.client.Set(ctx, key, val, 0).Err(); err != nil {
		logger.Error(ctx, "store_additional_fee_to_cache", "failed to set additional fee: %v", err)
		return err
	}

	return nil
}
