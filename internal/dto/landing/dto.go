package landing

type LandingRequest struct {
	Code string `json:"code" validate:"required"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Priority    int32  `json:"priority"`
}

type ItemCategory struct {
	Name     string `json:"name"`
	Priority int32  `json:"priority"`
	Items    []Item `json:"items"`
}

type LandingResponseData struct {
	Code           string         `json:"code"`
	Name           string         `json:"name"`
	ItemCategories []ItemCategory `json:"item_categories"`
}
