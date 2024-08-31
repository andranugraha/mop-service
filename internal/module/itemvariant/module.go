package itemvariant

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetActiveItemVariantsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariant, error)
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
