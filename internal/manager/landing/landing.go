package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemcategory"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
)

func (m *impl) Landing(ctx context.Context, code string) (*landing.LandingResponseData, error) {
	merchantData, err := m.merchantModule.GetMerchantOverview(ctx, code)
	if err != nil {
		return nil, err
	}

	if merchantData == nil {
		return nil, constant.ErrItemNotFound
	}

	return &landing.LandingResponseData{
		Code:           merchantData.GetCode(),
		Name:           merchantData.GetName(),
		ItemCategories: mapItemCategories(merchantData.GetItemCategories()),
	}, nil
}

func mapItemCategories(itemCategories []*itemcategory.ItemCategory) []landing.ItemCategory {
	mappedItemCategories := make([]landing.ItemCategory, len(itemCategories))
	for i, itemCategory := range itemCategories {
		mappedItemCategories[i] = landing.ItemCategory{
			Name:  itemCategory.GetName(),
			Items: mapItems(itemCategory.GetItems()),
		}
	}
	return mappedItemCategories
}

func mapItems(items []*item.Item) []landing.Item {
	mappedItems := make([]landing.Item, len(items))
	for i, item := range items {
		mappedItems[i] = landing.Item{
			Name:         item.GetName(),
			Description:  item.GetDescription(),
			Price:        item.GetPrice(),
			Priority:     item.GetPriority(),
			ItemVariants: mapItemVariants(item.GetVariants()),
		}
	}
	return mappedItems
}

func mapItemVariants(variants []*itemvariant.ItemVariant) []landing.ItemVariant {
	mappedVariants := make([]landing.ItemVariant, len(variants))
	for i, variant := range variants {
		mappedVariants[i] = landing.ItemVariant{
			Name:               variant.GetName(),
			MinSelect:          variant.GetMinSelect(),
			MaxSelect:          variant.GetMaxSelect(),
			ItemVariantOptions: mapItemVariantOptions(variant.GetOptions()),
		}
	}
	return mappedVariants
}

func mapItemVariantOptions(options []*itemvariantoption.ItemVariantOption) []landing.ItemVariantOption {
	mappedOptions := make([]landing.ItemVariantOption, len(options))
	for i, option := range options {
		mappedOptions[i] = landing.ItemVariantOption{
			Name:  option.GetName(),
			Price: option.GetPrice(),
		}
	}
	return mappedOptions
}
