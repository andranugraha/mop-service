package item

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetActiveItem(ctx context.Context, id uint64) (*Item, error)
	GetActiveItemsByIDs(ctx context.Context, ids []uint64) ([]*Item, error)
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
