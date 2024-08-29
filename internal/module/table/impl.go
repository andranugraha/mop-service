package table

import (
	"context"
)

func (i *impl) GetTable(ctx context.Context, id uint64) (*Table, error) {
	return i.dbStore.GetTable(ctx, id)
}
