package invoice

import (
	"context"
	"errors"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
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
