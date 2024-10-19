package banner

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetActiveBannersByMerchantID(ctx context.Context, merchantID uint64) ([]*Banner, error)
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
