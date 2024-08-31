package order

type CreateOrderRequest struct {
	MerchantID    uint64 `json:"merchant_id" validate:"required"`
	TableID       uint64 `json:"table_id" validate:"required"`
	Items         []Item `json:"items" validate:"required"`
	PaymentMethod uint32 `json:"payment_method"`
	TotalPrice    uint64 `json:"total_price" validate:"required"`
}

type CreateOrderResponse struct {
	OrderID   uint64 `json:"order_id"`
	OrderCode string `json:"order_code"`
	Total     uint64 `json:"total"`
	PaymentQR string `json:"payment_qr"`
	DueTime   uint64 `json:"due_time"`
}

type Item struct {
	ItemID   uint64         `json:"item_id" validate:"required"`
	Amount   uint64         `json:"amount" validate:"required"`
	Note     *string        `json:"note"`
	Variants []*ItemVariant `json:"variants"`
}

type ItemVariant struct {
	VariantID uint64   `json:"variant_id"`
	OptionIDs []uint64 `json:"option_ids"`
}
