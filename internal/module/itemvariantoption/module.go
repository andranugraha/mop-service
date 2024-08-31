package itemvariantoption

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetActiveItemVariantOptionsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariantOption, error)
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
