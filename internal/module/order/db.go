package order

import (
	"context"
	"errors"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/module/invoice"
	"gorm.io/gorm"
)

type db struct {
	client        *gorm.DB
	invoiceModule invoice.Module
}

func (d *db) CreateOrder(ctx context.Context, order *Order) error {
	return d.client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			logger.Error(ctx, "create_order", "failed to create order: %v", err.Error())
			return err
		}
		return nil
	})
}

func (d *db) GetFullOrderDataByID(ctx context.Context, id uint64) (*Order, error) {
	var order Order
	err := d.client.
		Preload("Invoice", func(db *gorm.DB) *gorm.DB {
			return db.Where("dtime is null")
		}).
		Where("id = ?", id).
		Where("dtime is null").
		Take(&order).Error
	if err != nil {
		logger.Error(ctx, "fetch_order_from_db", "failed to fetch order: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

func (d *db) UpdateOrder(ctx context.Context, order *Order) error {
	return d.client.Transaction(func(tx *gorm.DB) error {
		invoiceData := order.Invoice
		order.Invoice = nil
		err := tx.Save(order).Error
		if err != nil {
			tx.Rollback()
			logger.Error(ctx, "update_order", "failed to update order: %v", err.Error())
			return err
		}

		if invoiceData != nil {
			err = d.invoiceModule.UpdateInvoiceTx(ctx, tx, invoiceData)
			if err != nil {
				tx.Rollback()
				logger.Error(ctx, "update_order", "failed to update invoice: %v", err.Error())
				return err
			}
		}

		order.Invoice = invoiceData
		return nil
	})
}

func (d *db) GetOrderByInvoiceID(ctx context.Context, invoiceID uint64) (*Order, error) {
	var order Order
	err := d.client.
		Where("invoice_id = ?", invoiceID).
		Where("dtime is null").
		Take(&order).Error
	if err != nil {
		logger.Error(ctx, "fetch_order_from_db", "failed to fetch order: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}
