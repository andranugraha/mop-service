package itemcategory

import "context"

func (i *impl) GetItemCategoriesByMerchantId(
	ctx context.Context,
	merchantId uint64,
) ([]*ItemCategory, error) {
	itemCategories, err := i.cacheStore.GetItemCategoriesByMerchantId(ctx, merchantId)
	if err != nil {
		return nil, err
	}

	if len(itemCategories) > 0 {
		return itemCategories, nil
	}

	itemCategories, err = i.dbStore.GetItemCategoriesByMerchantId(ctx, merchantId)
	if err != nil {
		return nil, err
	}

	err = i.cacheStore.SetItemCategoriesByMerchantId(ctx, merchantId, itemCategories)
	if err != nil {
		return nil, err
	}

	return itemCategories, nil
}
