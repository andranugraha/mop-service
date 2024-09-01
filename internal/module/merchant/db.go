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
		Preload("ItemCategories", "dtime IS NULL").
		Preload("ItemCategories.Items", "dtime IS NULL").
		Preload("ItemCategories.Items.Variants", "dtime IS NULL").
		Preload("ItemCategories.Items.Variants.Options", "dtime IS NULL").
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

func (d *db) GetMerchantByID(ctx context.Context, id uint64) (*Merchant, error) {
	merchant := &Merchant{}
	err := d.client.
		Preload("PaymentTypes", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, merchant_id, name, type, extra_data").Where("dtime is null")
		}).
		Preload("AdditionalFees", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, merchant_id, name, type, description, fee").Where("dtime is null")
		}).
		Select("id, code, name").
		Where("id = ?", id).
		Where("dtime is null").
		Take(merchant).Error
	if err != nil {
		logger.Error(ctx, "fetch_merchant_from_db", "failed to fetch merchant: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return merchant, nil
}
