package invoice

import (
	"context"
	"errors"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
	"time"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetInvoiceByID(ctx context.Context, id uint64) (*Invoice, error) {
	var invoice Invoice
	err := d.client.
		Where("id = ?", id).
		Where("dtime is null").
		Take(&invoice).Error
	if err != nil {
		logger.Error(ctx, "fetch_invoice_from_db", "failed to fetch invoice: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (d *db) UpdateInvoiceTx(ctx context.Context, tx *gorm.DB, invoice *Invoice) error {
	if err := tx.Save(invoice).Error; err != nil {
		return err
	}
	return nil
}

func (d *db) GetTodayLatestInvoice(ctx context.Context, merchantID uint64) (*Invoice, error) {
	var invoice Invoice
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	err := d.client.
		Where("merchant_id = ?", merchantID).
		Where("status != ?", StatusCancelled).
		Where("dtime is null").
		Where("ctime >= ?", startOfDay.Unix()).
		Last(&invoice).Error
	if err != nil {
		logger.Error(ctx, "fetch_invoice_from_db", "failed to fetch invoice: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}
