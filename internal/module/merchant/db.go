package merchant

import (
	"context"
	"errors"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetMerchantByCode(ctx context.Context, code string) (*Merchant, error) {
	var merchant Merchant
	err := d.client.
		Where("code = ?", code).
		Take(&merchant).
		Error
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_db", "failed to fetch merchant: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &merchant, nil
}
