package itemvariant

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

func (c *cache) getItemVariantKey(id uint64) string {
	return fmt.Sprintf("item_variant:%d", id)
}

func (c *cache) GetActiveItemVariantsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariant, []uint64, error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, c.getItemVariantKey(id))
	}

	val, err := c.client.MGet(ctx, keys...).Result()
	if err != nil {
		logger.Error(ctx, "fetch_item_variants_from_cache", "failed to fetch item variants: %v", err)
		return nil, ids, err
	}

	itemVariants := make([]*ItemVariant, 0, len(ids))
	missedIDs := make([]uint64, 0, len(ids))
	for i, v := range val {
		if v == nil {
			missedIDs = append(missedIDs, ids[i])
			continue
		}
		var itemVariant ItemVariant
		err = json.Unmarshal([]byte(v.(string)), &itemVariant)
		if err != nil {
			logger.Error(ctx, "fetch_item_variants_from_cache", "failed to unmarshal item variant: %v", err)
			missedIDs = append(missedIDs, ids[i])
		}
		itemVariants = append(itemVariants, &itemVariant)
	}

	return itemVariants, missedIDs, nil
}

func (c *cache) SetManyActiveItemVariants(ctx context.Context, itemVariants []*ItemVariant) error {
	pipe := c.client.Pipeline()
	for _, itemVariant := range itemVariants {
		key := c.getItemVariantKey(itemVariant.GetId())
		jsonValue, err := json.Marshal(itemVariant)
		if err != nil {
			logger.Error(ctx, "set_item_variant_to_cache", "failed to marshal item variant: %v", err)
			return err
		}
		pipe.Set(ctx, key, jsonValue, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error(ctx, "set_item_variants_to_cache", "failed to set item variants: %v", err)
		return err
	}

	return nil
}
