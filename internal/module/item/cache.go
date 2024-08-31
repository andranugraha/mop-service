package item

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

func (c *cache) getItemKey(id uint64) string {
	return fmt.Sprintf("item:%d", id)
}

func (c *cache) GetActiveItem(ctx context.Context, id uint64) (*Item, error) {
	key := c.getItemKey(id)
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_item_from_cache", "failed to get item: %v", err)
		return nil, err
	}

	var item Item
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		logger.Error(ctx, "fetch_item_from_cache", "failed to unmarshal item: %v", err)
		return nil, err
	}

	return &item, nil
}

func (c *cache) GetActiveItemsByIDs(ctx context.Context, ids []uint64) ([]*Item, []uint64, error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, c.getItemKey(id))
	}

	val, err := c.client.MGet(ctx, keys...).Result()
	if err != nil {
		logger.Error(ctx, "fetch_items_from_cache", "failed to get items: %v", err)
		return nil, ids, err
	}

	items := make([]*Item, 0, len(ids))
	missedIDs := make([]uint64, 0, len(ids))
	for i, v := range val {
		if v == nil {
			missedIDs = append(missedIDs, ids[i])
			continue
		}
		var item Item
		err = json.Unmarshal([]byte(v.(string)), &item)
		if err != nil {
			logger.Error(ctx, "fetch_items_from_cache", "failed to unmarshal item: %v", err)
			missedIDs = append(missedIDs, ids[i])
		}
		items = append(items, &item)
	}

	return items, missedIDs, nil
}

func (c *cache) SetActiveItem(ctx context.Context, item *Item) error {
	key := c.getItemKey(item.GetId())
	jsonValue, err := json.Marshal(item)
	if err != nil {
		logger.Error(ctx, "set_item_to_cache", "failed to marshal item: %v", err)
		return err
	}

	err = c.client.Set(ctx, key, jsonValue, 0).Err()
	if err != nil {
		logger.Error(ctx, "set_item_to_cache", "failed to set item: %v", err)
		return err
	}

	return nil
}

func (c *cache) SetManyActiveItems(ctx context.Context, items []*Item) error {
	pipe := c.client.Pipeline()
	for _, item := range items {
		key := c.getItemKey(item.GetId())
		jsonValue, err := json.Marshal(item)
		if err != nil {
			logger.Error(ctx, "set_items_to_cache", "failed to marshal item: %v", err)
			return err
		}
		pipe.Set(ctx, key, jsonValue, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error(ctx, "set_items_to_cache", "failed to set items: %v", err)
		return err
	}

	return nil
}
