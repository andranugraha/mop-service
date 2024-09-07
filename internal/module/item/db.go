package item

import (
	"context"
	"errors"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetItem(ctx context.Context, id uint64) (*Item, error) {
	item := &Item{}
	err := d.client.
		Preload("Variants", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Select("id, name, description, price").
		Where("id = ?", id).
		Where("dtime is null").
		Take(item).Error
	if err != nil {
		logger.Error(ctx, "fetch_item_from_db", "failed to fetch item: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func (d *db) GetActiveItemsByIDs(ctx context.Context, ids []uint64) ([]*Item, error) {
	var items []*Item
	err := d.client.
		Select("id, name, description, price").
		Where("id IN (?)", ids).
		Where("dtime is null").
		Find(&items).Error
	if err != nil {
		logger.Error(ctx, "fetch_items_from_db", "failed to fetch items: %v", err.Error())
		return nil, err
	}

	return items, nil
}

func (d *db) GetActiveItemsByCategoryId(ctx context.Context, categoryID uint64) ([]*Item, error) {
	var items []*Item
	err := d.client.
		Where("item_category_id = ?", categoryID).
		Where("dtime is null").
		Order("priority DESC").
		Find(&items).Error
	if err != nil {
		logger.Error(ctx, "fetch_items_from_db", "failed to fetch items: %v", err.Error())
		return nil, err
	}

	return items, nil
}
