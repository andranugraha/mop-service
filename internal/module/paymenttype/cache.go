package paymenttype

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

func (c *cache) getMerchantPaymentTypeKey(merchantID uint64) string {
	return fmt.Sprintf("payment_type:merchant_id:%d", merchantID)
}

func (c *cache) GetActivePaymentTypesByMerchantID(ctx context.Context, merchantID uint64) ([]*PaymentType, error) {
	key := c.getMerchantPaymentTypeKey(merchantID)
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_payment_types_from_cache", "failed to get payment types: %v", err)
		return nil, err
	}

	var paymentTypes []*PaymentType
	err = json.Unmarshal([]byte(val), &paymentTypes)
	if err != nil {
		logger.Error(ctx, "fetch_payment_types_from_cache", "failed to unmarshal payment types: %v", err)
		return nil, err
	}

	return paymentTypes, nil
}

func (c *cache) SetPaymentTypesByMerchantID(ctx context.Context, merchantID uint64, paymentTypes []*PaymentType) error {
	key := c.getMerchantPaymentTypeKey(merchantID)
	val, err := json.Marshal(paymentTypes)
	if err != nil {
		logger.Error(ctx, "store_payment_types_to_cache", "failed to marshal payment types: %v", err)
		return err
	}

	if err = c.client.Set(ctx, key, val, 0).Err(); err != nil {
		logger.Error(ctx, "store_payment_types_to_cache", "failed to set payment types: %v", err)
		return err
	}

	return nil
}
