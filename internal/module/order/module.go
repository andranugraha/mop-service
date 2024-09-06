package order

import (
	"context"
	"github.com/empnefsi/mop-service/internal/module/invoice"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetFullOrderDataByID(ctx context.Context, id uint64) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
}

type impl struct {
	dbStore *db
}

func GetModule() Module {
	return &impl{
		dbStore: &db{
			client:        config.GetDB(),
			invoiceModule: invoice.GetModule(),
		},
	}
}
