package banner

import (
	"context"
	"errors"
	"time"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetActiveBannersByMerchantID(ctx context.Context, merchantID uint64) ([]*Banner, error) {
	var banners []*Banner
	now := time.Now().Unix()
	err := d.client.
		Where("merchant_id = ?", merchantID).
		Where("(end_date > ? or end_date is null)", now).
		Where("dtime is null").
		Order("priority ASC").
		Find(&banners).Error
	if err != nil {
		logger.Error(ctx, "fetch_active_banners_from_db", "failed to fetch banners: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	logger.InfoWithData(ctx, "fetch_active_banners_from_db", "banners: %v", banners)
	return banners, nil
}
