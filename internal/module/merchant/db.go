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

func (d *db) GetMerchantOverview(ctx context.Context, code string) (*Merchant, error) {
	var merchant Merchant
	err := d.client.
		Preload("ItemCategories", func(db *gorm.DB) *gorm.DB {
			return db.Where("dtime IS NULL").Order("priority ASC")
		}).
		Preload("ItemCategories.Items", func(db *gorm.DB) *gorm.DB {
			return db.Where("dtime IS NULL").Order("priority ASC")
		}).
		Preload("ItemCategories.Items.Variants", "dtime IS NULL").
		Preload("ItemCategories.Items.Variants.Options", "dtime IS NULL").
		Where("code = ?", code).
		Where("dtime is null").
		Take(&merchant).
		Error

	if err != nil {
		logger.Error(
			ctx, "fetch_merchant_overview_from_db", "failed to fetch merchant: %v", err.Error(),
		)
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
