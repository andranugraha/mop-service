package table

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetTableByID(ctx context.Context, id uint64) (*Table, error)
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
