package order

import (
	"context"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) CreateOrder(ctx context.Context, order *Order) error {
	return d.client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})
}
