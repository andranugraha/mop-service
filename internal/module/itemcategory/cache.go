package itemcategory

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

func (c *cache) getItemCategoryKey(id uint64) string {
	return fmt.Sprintf("item_category:merchant_id:%d", id)
}

func (c *cache) SetItemCategoriesByMerchantId(
	ctx context.Context, merchantId uint64, itemCategories []*ItemCategory,
) error {
	key := c.getItemCategoryKey(merchantId)

	pipe := c.client.Pipeline()
	for _, itemCategory := range itemCategories {
		jsonValue, err := json.Marshal(itemCategory)
		if err != nil {
			logger.Error(ctx, "set_item_categories_to_cache",
				"failed to marshal item category: %v", err)
			return err
		}

		score := float64(*itemCategory.Priority)
		pipe.ZAdd(ctx, key, &redis.Z{
			Score:  score,
			Member: jsonValue,
		})
	}

	expiryInSeconds := config.GetCacheUserExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	pipe.Expire(ctx, key, expiryDuration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error(ctx, "set_item_categories_to_cache",
			"failed to set item categories: %v", err)
		return err
	}

	logger.Data(ctx, "set_item_categories_to_cache",
		"item categories set to cache: %v", itemCategories)
	return nil
}

func (c *cache) GetItemCategoriesByMerchantId(
	ctx context.Context, merchantId uint64,
) ([]*ItemCategory, error) {
	key := c.getItemCategoryKey(merchantId)
	jsonValues, err := c.client.ZRevRange(ctx, key, 0, -1).Result()

	if err != nil {
		logger.Error(ctx, "fetch_item_categories_from_cache",
			"failed to fetch item categories: %v", err)
		return nil, err
	}

	itemCategories := make([]*ItemCategory, 0, len(jsonValues))
	for _, jsonValue := range jsonValues {
		itemCategory := &ItemCategory{}
		err = json.Unmarshal([]byte(jsonValue), itemCategory)
		if err != nil {
			logger.Warn(ctx, "fetch_item_categories_from_cache",
				"failed to unmarshal item category: %v", err)
			continue
		}

		itemCategories = append(itemCategories, itemCategory)
	}

	logger.Data(ctx, "fetch_item_categories_from_cache",
		"item categories fetched from cache: %v", itemCategories)
	return itemCategories, nil
}
