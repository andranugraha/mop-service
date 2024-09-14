package merchant

type GetMerchantActivePaymentTypesRequest struct {
	MerchantID uint64 `json:"merchant_id"`
}

type GetMerchantActivePaymentTypesResponse struct {
	PaymentTypes []*PaymentType `json:"payment_types"`
}

type PaymentType struct {
	Type uint32 `json:"type"`
	Name string `json:"name"`
}
