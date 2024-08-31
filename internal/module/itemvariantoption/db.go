package itemvariantoption

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetActiveItemVariantsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariantOption, error) {
	var itemVariantOptions []*ItemVariantOption
	err := d.client.
		Select("id, name, item_variant_id, price").
		Where("id IN ?", ids).
		Where("dtime is null").
		Find(&itemVariantOptions).Error
	if err != nil {
		logger.Error(ctx, "get_item_variant_options_by_ids", "failed to fetch item variant options: %v", err)
		return nil, err
	}

	return itemVariantOptions, nil
}
