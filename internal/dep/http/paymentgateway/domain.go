package paymentgateway

import (
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"google.golang.org/protobuf/proto"
)

const (
	PaymentEndpoint       = "/charge"
	CancelPaymentEndpoint = "/cancel"
)

const (
	PaymentTypeQRIS = "qris"
)

const (
	ActionGenerateQRCode = "generate-qr-code"
)

var paymentType = map[uint32]string{
	paymenttype.PaymentTypeQR: PaymentTypeQRIS,
}

const (
	UnitDay    = "day"
	UnitHour   = "hour"
	UnitMinute = "minute"
	UnitSecond = "second"
)

const (
	TransactionStatusSettlement = "settlement"
)

type PaymentRequest struct {
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	ItemDetails        []*ItemDetail      `json:"item_details"`
	CustomerDetails    *CustomerDetail    `json:"customer_details"`
	CustomExpiry       *CustomExpiry      `json:"custom_expiry"`
	QRIS               *QRISDetail        `json:"qris"`
}

type CancelPaymentRequest struct {
	OrderID string `json:"order_id"`
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type ItemDetail struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
}

type CustomerDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CustomExpiry struct {
	OrderTime      string `json:"order_time"`
	ExpiryDuration int    `json:"expiry_duration"`
	Unit           string `json:"unit"`
}

type QRISDetail struct {
	Acquirer string `json:"acquirer"`
}

type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

type PaymentResponse struct {
	StatusCode        string   `json:"status_code"`
	StatusMessage     string   `json:"status_message"`
	TransactionID     string   `json:"transaction_id"`
	OrderID           string   `json:"order_id"`
	MerchantID        string   `json:"merchant_id"`
	GrossAmount       string   `json:"gross_amount"`
	Currency          string   `json:"currency"`
	PaymentType       string   `json:"payment_type"`
	TransactionTime   string   `json:"transaction_time"`
	TransactionStatus string   `json:"transaction_status"`
	FraudStatus       string   `json:"fraud_status"`
	Actions           []Action `json:"actions"`
	QRString          string   `json:"qr_string"`
	Acquirer          string   `json:"acquirer"`
	ExpiryTime        string   `json:"expiry_time"`
}

type CancelPaymentResponse struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	TransactionID string `json:"transaction_id"`
	OrderID       string `json:"order_id"`
}

func GetPaymentType(intPaymentType uint32) string {
	return paymentType[intPaymentType]
}

func (p *PaymentResponse) GetQRURL() *string {
	if p == nil {
		return nil
	}

	for _, action := range p.Actions {
		if action.Name == ActionGenerateQRCode {
			return proto.String(action.URL)
		}
	}

	return nil
}
