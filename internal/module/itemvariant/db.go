package itemvariant

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetActiveItemVariantsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariant, error) {
	var itemVariants []*ItemVariant
	err := d.client.
		Select("id, name, item_id, min_select, max_select").
		Where("id IN (?)", ids).
		Where("dtime is null").
		Find(&itemVariants).Error
	if err != nil {
		logger.Error(ctx, "get_active_item_variants_by_ids", "failed to fetch item variants: %v", err)
		return nil, err
	}

	return itemVariants, nil
}
