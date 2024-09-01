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

	itemCategories := make([]landing.ItemCategory, len(merchantData.GetItemCategories()))
	for i, category := range merchantData.GetItemCategories() {
		items := make([]landing.Item, len(category.GetItems()))

		for j, item := range category.GetItems() {
			itemVariants := make([]landing.ItemVariant, len(item.GetVariants()))

			for k, variant := range item.GetVariants() {
				itemVariantOptions := make([]landing.ItemVariantOption, len(variant.GetOptions()))

				for l, option := range variant.GetOptions() {
					itemVariantOptions[l] = landing.ItemVariantOption{
						Name:  option.GetName(),
						Price: option.GetPrice(),
					}
				}
				itemVariants[k] = landing.ItemVariant{
					Name:               variant.GetName(),
					MinSelect:          variant.GetMinSelect(),
					MaxSelect:          variant.GetMaxSelect(),
					ItemVariantOptions: itemVariantOptions,
				}
			}
			items[j] = landing.Item{
				Name:         item.GetName(),
				Description:  item.GetDescription(),
				Price:        item.GetPrice(),
				Priority:     item.GetPriority(),
				ItemVariants: itemVariants,
			}
		}
		itemCategories[i] = landing.ItemCategory{
			Name:     category.GetName(),
			Priority: category.GetPriority(),
			Items:    items,
		}
	}

	return &landing.LandingResponseData{
		Code:           merchantData.GetCode(),
		Name:           merchantData.GetName(),
		ItemCategories: itemCategories,
	}, nil
}
