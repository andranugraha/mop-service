package merchant

type MerchantRequest struct {
	Code string `json:"code" validate:"required"`
}

type MerchantResponse struct {
	Data MerchantResponseData `json:"data"`
}

type MerchantResponseData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
