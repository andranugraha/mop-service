package additionalfee

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetActiveAdditionalFeesByMerchantID(ctx context.Context, merchantID uint64) ([]*AdditionalFee, error)
}

type impl struct {
	cacheStore *cache
	dbStore    *db
}

func GetModule() Module {
	return &impl{
		cacheStore: &cache{
			client: config.GetCache(),
		},
		dbStore: &db{
			client: config.GetDB(),
		},
	}
}
