package itemcategory

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetItemCategoriesByMerchantId(
	ctx context.Context,
	merchantID uint64,
) ([]*ItemCategory, error) {
	var itemCategories []*ItemCategory
	err := d.client.
		Where("merchant_id = ?", merchantID).
		Where("dtime IS NULL").
		Order("priority DESC").
		Find(&itemCategories).
		Error
	if err != nil {
		logger.Error(ctx, "fetch_item_categories_from_db",
			"failed to fetch item categories: %v", err.Error())
		return nil, err
	}

	return itemCategories, nil
}
