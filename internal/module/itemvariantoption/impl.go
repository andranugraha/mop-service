package itemvariantoption

import (
	"context"
)

func (i *impl) GetActiveItemVariantOptionsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariantOption, error) {
	itemVariantOptions, missedIDs, _ := i.cacheStore.GetActiveItemVariantOptionsByIDs(ctx, ids)
	if len(missedIDs) > 0 {
		dbItemVariantOptions, err := i.dbStore.GetActiveItemVariantsByIDs(ctx, missedIDs)
		if err != nil {
			return nil, err
		}

		for _, itemVariantOption := range dbItemVariantOptions {
			itemVariantOptions = append(itemVariantOptions, itemVariantOption)
		}

		if len(dbItemVariantOptions) > 0 {
			go func() {
				_ = i.cacheStore.SetManyItemVariantOptions(ctx, dbItemVariantOptions)
			}()
		}
	}

	return itemVariantOptions, nil
}
