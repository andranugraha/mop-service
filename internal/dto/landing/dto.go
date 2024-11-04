package landing

type LandingRequest struct {
	Code string `json:"code" validate:"required"`
}

type ItemCategory struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Priority int32  `json:"priority"`
	Icon     string `json:"icon"`
	Items    []Item `json:"items"`
}

type Item struct {
	Id            uint64        `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Images        []string      `json:"images"`
	Price         uint64        `json:"price"`
	Priority      int32         `json:"priority"`
	IsRecommended bool          `json:"is_recommended"`
	ItemVariants  []ItemVariant `json:"item_variants"`
}

type ItemVariant struct {
	Id                 uint64              `json:"id"`
	Name               string              `json:"name"`
	MinSelect          uint32              `json:"min_select"`
	MaxSelect          uint32              `json:"max_select"`
	ItemVariantOptions []ItemVariantOption `json:"item_variant_options"`
}

type ItemVariantOption struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type LandingResponse struct {
	Id             uint64         `json:"id"`
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	ItemCategories []ItemCategory `json:"item_categories"`
}

type GetLandingBannersResponse struct {
	Banners []Banner `json:"banners"`
}

type Banner struct {
	Id          uint64  `json:"id"`
	Image       string  `json:"image"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartDate   uint64  `json:"start_date"`
	EndDate     *uint64 `json:"end_date"`
}
