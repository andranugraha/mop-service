package order

import "mime/multipart"

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

type PayOrderRequest struct {
	OrderID        uint64                `form:"order_id" validate:"required"`
	ProofOfPayment *multipart.FileHeader `form:"proof_of_payment" validate:"required"`
}

type PayOrderResponse struct {
	InvoiceID   uint64 `json:"invoice_id"`
	InvoiceCode string `json:"invoice_code"`
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

type PaymentCallbackRequest struct {
	TransactionType   string `json:"transaction_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	SettlementTime    string `json:"settlement_time"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	Issuer            string `json:"issuer"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Currency          string `json:"currency"`
	Acquirer          string `json:"acquirer"`
}
