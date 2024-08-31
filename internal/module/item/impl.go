package item

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (i *impl) GetActiveItem(ctx context.Context, id uint64) (*Item, error) {
	item, _ := i.cacheStore.GetActiveItem(ctx, id)
	if item != nil {
		return item, nil
	}

	item, err := i.dbStore.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, constant.ErrItemNotFound
	}

	go func() {
		_ = i.cacheStore.SetActiveItem(ctx, item)
	}()
	return item, nil
}

func (i *impl) GetActiveItemsByIDs(ctx context.Context, ids []uint64) ([]*Item, error) {
	items, missedIDs, _ := i.cacheStore.GetActiveItemsByIDs(ctx, ids)
	if len(missedIDs) == 0 {
		return items, nil
	}

	dbItems, err := i.dbStore.GetActiveItemsByIDs(ctx, missedIDs)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = i.cacheStore.SetManyActiveItems(ctx, dbItems)
	}()

	items = append(items, dbItems...)
	return items, nil
}
