package landing

type LandingRequest struct {
	Code string `json:"code" validate:"required"`
}

type ItemCategory struct {
	Name     string `json:"name"`
	Priority int32  `json:"priority"`
	Items    []Item `json:"items"`
}

type Item struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Price        uint64        `json:"price"`
	Priority     int32         `json:"priority"`
	ItemVariants []ItemVariant `json:"item_variants"`
}

type ItemVariant struct {
	Name               string              `json:"name"`
	MinSelect          uint32              `json:"min_select"`
	MaxSelect          uint32              `json:"max_select"`
	ItemVariantOptions []ItemVariantOption `json:"item_variant_options"`
}

type ItemVariantOption struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type LandingResponseData struct {
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	ItemCategories []ItemCategory `json:"item_categories"`
}
