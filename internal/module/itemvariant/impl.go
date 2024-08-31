package itemvariant

import "context"

func (i *impl) GetActiveItemVariantsByIDs(ctx context.Context, ids []uint64) ([]*ItemVariant, error) {
	itemVariants, missedIDs, _ := i.cacheStore.GetActiveItemVariantsByIDs(ctx, ids)
	if len(missedIDs) > 0 {
		dbItemVariants, err := i.dbStore.GetActiveItemVariantsByIDs(ctx, missedIDs)
		if err != nil {
			return nil, err
		}

		for _, itemVariant := range dbItemVariants {
			itemVariants = append(itemVariants, itemVariant)
		}

		if len(dbItemVariants) > 0 {
			go func() {
				_ = i.cacheStore.SetManyActiveItemVariants(ctx, dbItemVariants)
			}()
		}
	}

	return itemVariants, nil
}
