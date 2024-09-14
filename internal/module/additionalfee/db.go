package additionalfee

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetActiveAdditionalFeesByMerchantID(ctx context.Context, merchantID uint64) ([]*AdditionalFee, error) {
	var additionalFee []*AdditionalFee
	err := d.client.
		Where("merchant_id = ? AND dtime is null", merchantID).
		Find(&additionalFee).Error
	if err != nil {
		logger.Error(ctx, "fetch_additional_fee_from_db", "failed to get additional fee: %v", err)
		return nil, err
	}

	return additionalFee, nil
}
