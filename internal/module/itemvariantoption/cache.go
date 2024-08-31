package itemvariantoption

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

func (c *cache) getItemVariantOptionKey(id uint64) string {
	return fmt.Sprintf("item_variant_option:%d", id)
}

func (c *cache) GetActiveItemVariantOptionsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariantOption, []uint64, error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, c.getItemVariantOptionKey(id))
	}

	val, err := c.client.MGet(ctx, keys...).Result()
	if err != nil {
		logger.Error(ctx, "fetch_item_variant_options_from_cache", "failed to get item variant options: %v", err)
		return nil, ids, err
	}

	itemVariantOptions := make([]*ItemVariantOption, 0, len(ids))
	missedIDs := make([]uint64, 0, len(ids))
	for i, v := range val {
		if v == nil {
			missedIDs = append(missedIDs, ids[i])
			continue
		}
		var itemVariantOption ItemVariantOption
		err = json.Unmarshal([]byte(v.(string)), &itemVariantOption)
		if err != nil {
			logger.Error(ctx, "fetch_item_variant_options_from_cache", "failed to unmarshal item variant option: %v", err)
			missedIDs = append(missedIDs, ids[i])
		}
		itemVariantOptions = append(itemVariantOptions, &itemVariantOption)
	}

	return itemVariantOptions, missedIDs, nil
}

func (c *cache) SetManyItemVariantOptions(ctx context.Context, itemVariantOptions []*ItemVariantOption) error {
	pipe := c.client.Pipeline()
	for _, itemVariantOption := range itemVariantOptions {
		key := c.getItemVariantOptionKey(itemVariantOption.GetId())
		jsonValue, err := json.Marshal(itemVariantOption)
		if err != nil {
			logger.Error(ctx, "set_item_variant_options_to_cache", "failed to marshal item variant option: %v", err)
			continue
		}
		pipe.Set(ctx, key, jsonValue, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error(ctx, "set_item_variant_options_to_cache", "failed to set item variant options: %v", err)
	}

	return nil
}
