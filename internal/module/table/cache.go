package table

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

func (c *cache) getTableKeyPatternByCode(code string) string {
	return fmt.Sprintf("table_data:tableid:*:code:%s", code)
}

func (c *cache) getTableKeyPatternByID(id uint64) string {
	return fmt.Sprintf("table_data:tableid:%d:*", id)
}

func (c *cache) getTableKey(id uint64, code string) string {
	return fmt.Sprintf("table_data:tableid:%d:code:%s", id, code)
}

func (c *cache) GetTableByID(ctx context.Context, id uint64) (*Table, error) {
	table := &Table{}
	keyPattern := c.getTableKeyPatternByID(id)
	keys, err := c.client.Keys(ctx, keyPattern).Result()
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", err)
		return nil, err
	}

	if len(keys) == 0 {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", errors.New("table not found"))
		return nil, nil
	}

	key := keys[0]
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &table)
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to unmarshal table: %v", err)
		return nil, err
	}

	return table, nil
}

func (c *cache) GetTableByCode(ctx context.Context, code string) (*Table, error) {
	table := &Table{}
	keyPattern := c.getTableKeyPatternByCode(code)
	keys, err := c.client.Keys(ctx, keyPattern).Result()
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", err)
		return nil, err
	}

	if len(keys) == 0 {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", errors.New("table not found"))
		return nil, nil
	}

	key := keys[0]
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to get table: %v", err)
		return nil, err
	}

	err = json.Unmarshal([]byte(val), &table)
	if err != nil {
		logger.Error(ctx, "fetch_table_from_cache", "failed to unmarshal table: %v", err)
		return nil, err
	}

	return table, nil
}

func (c *cache) SetTable(ctx context.Context, table *Table) error {
	key := c.getTableKey(table.GetId(), table.GetCode())
	jsonValue, err := json.Marshal(table)
	if err != nil {
		logger.Error(ctx, "set_table_to_cache", "failed to marshal table: %v", err)
		return err
	}

	expiryInSeconds := config.GetCacheTableExpiry()
	expiryDuration := time.Duration(expiryInSeconds) * time.Second
	err = c.client.Set(ctx, key, jsonValue, expiryDuration).Err()
	if err != nil {
		logger.Error(ctx, "set_table_to_cache", "failed to set table: %v", err)
		return err
	}

	return nil
}
