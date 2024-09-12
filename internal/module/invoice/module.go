package invoice

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
	"gorm.io/gorm"
)

type Module interface {
	GetInvoiceByID(ctx context.Context, id uint64) (*Invoice, error)
	UpdateInvoiceTx(ctx context.Context, tx *gorm.DB, invoice *Invoice) error
	GetTodayLatestInvoice(ctx context.Context, merchantID uint64) (*Invoice, error)
}

type impl struct {
	dbStore *db
}

func GetModule() Module {
	return &impl{
		dbStore: &db{
			client: config.GetDB(),
		},
	}
}
