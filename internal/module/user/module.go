package user

import (
	"context"

	"github.com/empnefsi/mop-service/internal/config"
)

type Module interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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
