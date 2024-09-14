package merchant

type GetMerchantActivePaymentTypesResponse struct {
	PaymentTypes []*PaymentType `json:"payment_types"`
}

type PaymentType struct {
	Type uint32 `json:"type"`
	Name string `json:"name"`
}

type GetMerchantActiveAdditionalFeesResponse struct {
	AdditionalFees []*AdditionalFee `json:"additional_fees"`
}

type AdditionalFee struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        uint32 `json:"type"`
	Fee         uint64 `json:"fee"`
}
