package table

import (
	"context"
	"errors"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetTable(ctx context.Context, id uint64) (*Table, error) {
	var data Table
	err := d.client.
		Select("id, code, merchant_id").
		Where("dtime is null").
		Where("id = ?", id).
		Take(&data).Error
	if err != nil {
		logger.Error(ctx, "fetch_table_from_db", "failed to fetch table: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
