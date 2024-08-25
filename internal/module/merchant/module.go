package merchant

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetMerchantByCode(ctx context.Context, code string) (*Merchant, error)
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
