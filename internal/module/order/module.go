package order

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	CreateOrder(ctx context.Context, order *Order) error
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
