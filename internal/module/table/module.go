package table

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetTable(ctx context.Context, id uint64) (*Table, error)
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
