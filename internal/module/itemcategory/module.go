package itemcategory

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetItemCategoriesByMerchantId(ctx context.Context, merchantId uint64) ([]*ItemCategory, error)
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
