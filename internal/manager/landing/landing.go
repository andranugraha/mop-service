package landing

import (
	"context"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemcategory"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
)

func (m *impl) Landing(ctx context.Context, code string) (*landing.LandingResponse, error) {
	merchantData, err := m.merchantModule.GetMerchantOverview(ctx, code)
	if err != nil {
		return nil, err
	}

	return &landing.LandingResponse{
		Id:             merchantData.GetId(),
		Code:           merchantData.GetCode(),
		Name:           merchantData.GetName(),
		ItemCategories: constructCategories(merchantData.GetItemCategories()),
	}, nil
}

func constructCategories(itemCategories []*itemcategory.ItemCategory) []landing.ItemCategory {
	categories := make([]landing.ItemCategory, len(itemCategories))
	for i, itemCategory := range itemCategories {
		categories[i] = landing.ItemCategory{
			Id:       itemCategory.GetId(),
			Name:     itemCategory.GetName(),
			Priority: itemCategory.GetPriority(),
			Icon:     config.GetCDNEndpoint() + "/" + itemCategory.GetIcon(),
			Items:    constructItems(itemCategory.GetItems()),
		}
	}
	return categories
}

func constructItems(items []*item.Item) []landing.Item {
	mappedItems := make([]landing.Item, len(items))
	for i, item := range items {
		mappedItems[i] = landing.Item{
			Id:           item.GetId(),
			Name:         item.GetName(),
			Description:  item.GetDescription(),
			Image:        config.GetCDNEndpoint() + "/" + item.GetImage(),
			Price:        item.GetPrice(),
			Priority:     item.GetPriority(),
			ItemVariants: constructItemVariants(item.GetVariants()),
		}
	}
	return mappedItems
}

func constructItemVariants(variants []*itemvariant.ItemVariant) []landing.ItemVariant {
	mappedVariants := make([]landing.ItemVariant, len(variants))
	for i, variant := range variants {
		mappedVariants[i] = landing.ItemVariant{
			Id:                 variant.GetId(),
			Name:               variant.GetName(),
			MinSelect:          variant.GetMinSelect(),
			MaxSelect:          variant.GetMaxSelect(),
			ItemVariantOptions: constructItemVariantOptions(variant.GetOptions()),
		}
	}
	return mappedVariants
}

func constructItemVariantOptions(options []*itemvariantoption.ItemVariantOption) []landing.ItemVariantOption {
	mappedOptions := make([]landing.ItemVariantOption, len(options))
	for i, option := range options {
		mappedOptions[i] = landing.ItemVariantOption{
			Id:    option.GetId(),
			Name:  option.GetName(),
			Price: option.GetPrice(),
		}
	}
	return mappedOptions
}
