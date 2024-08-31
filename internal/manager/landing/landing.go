package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/dto/landing"
)

func (m *impl) Landing(ctx context.Context, code string) (*landing.LandingResponseData, error) {
	merchantData, err := m.merchantModule.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if merchantData == nil {
		return nil, constant.ErrItemNotFound
	}

	itemCategoriesData, err := m.itemcategoryModule.GetItemCategoriesByMerchantId(ctx, merchantData.GetId())
	if err != nil {
		return nil, err
	}

	itemCategories := make([]landing.ItemCategory, len(itemCategoriesData))
	for i, itemCategoryData := range itemCategoriesData {
		itemCategories[i] = landing.ItemCategory{
			Name:     itemCategoryData.GetName(),
			Priority: itemCategoryData.GetPriority(),
		}

		itemsData, err := m.itemModule.GetActiveItemsByCategoryId(ctx, itemCategoryData.GetId())
		if err != nil {
			return nil, err
		}

		items := make([]landing.Item, len(itemsData))
		for j, itemData := range itemsData {
			items[j] = landing.Item{
				Name:        itemData.GetName(),
				Description: itemData.GetDescription(),
				Price:       itemData.GetPrice(),
				Priority:    itemData.GetPriority(),
			}
		}

		itemCategories[i].Items = items
	}

	return &landing.LandingResponseData{
		Code:           merchantData.GetCode(),
		Name:           merchantData.GetName(),
		ItemCategories: itemCategories,
	}, nil
}
