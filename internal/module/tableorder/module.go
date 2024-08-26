package tableorder

import "github.com/empnefsi/mop-service/internal/config"

type Module interface {
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
