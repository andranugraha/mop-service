package table

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (i *impl) GetTableByID(ctx context.Context, id uint64) (*Table, error) {
	table, _ := i.cacheStore.GetTableByID(ctx, id)
	if table != nil {
		return table, nil
	}

	table, err := i.dbStore.GetTableByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if table == nil {
		return nil, constant.ErrTableNotFound
	}

	go func() {
		_ = i.cacheStore.SetTable(ctx, table)
	}()

	return table, nil
}
