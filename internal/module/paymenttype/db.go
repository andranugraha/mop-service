package paymenttype

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetActivePaymentTypesByMerchantID(ctx context.Context, merchantID uint64) ([]*PaymentType, error) {
	var paymentTypes []*PaymentType
	if err := d.client.
		Where("merchant_id = ?", merchantID).
		Where("dtime is null").
		Find(&paymentTypes).Error; err != nil {
		logger.Error(ctx, "fetch_payment_types_from_db", "failed to get payment types: %v", err)
		return nil, err
	}
	return paymentTypes, nil
}
