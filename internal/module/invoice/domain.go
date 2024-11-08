package invoice

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"gorm.io/gorm"
)

const tableName = "invoice_tab"

const (
	StatusPending uint32 = iota
	StatusPaid
	StatusCancelled
)

type Invoice struct {
	Id             *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId     *uint64 `json:"merchant_id"`
	PaymentTypeId  *uint64 `json:"payment_type_id"`
	Code           *string `json:"code"`
	AdditionalFees []byte  `json:"additional_fees"`
	TotalPayment   *uint64 `json:"total_payment"`
	PaymentProof   *string `json:"payment_proof"`
	Status         *uint32 `json:"status"`
	Ctime          *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime          *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime          *uint64 `json:"dtime"`

	PaymentType *paymenttype.PaymentType `gorm:"foreignKey:PaymentTypeId;references:Id" json:"payment_type"`
	Merchant    *merchant.Merchant       `gorm:"foreignKey:MerchantId;references:Id" json:"merchant"`
}

type AdditionalFee struct {
	Id     *uint64 `json:"id"`
	Type   *uint32 `json:"type"`
	Name   *string `json:"name"`
	Fee    *uint64 `json:"fee"`
	Amount *uint64 `json:"amount"`
}

func (i *Invoice) TableName() string {
	return tableName
}

func (i *Invoice) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *Invoice) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *Invoice) GetPaymentTypeId() uint64 {
	if i.PaymentTypeId != nil {
		return *i.PaymentTypeId
	}
	return 0
}

func (i *Invoice) GetCode() string {
	if i.Code != nil {
		return *i.Code
	}
	return ""
}

func (i *Invoice) GetAdditionalFees() []*AdditionalFee {
	var additionalFees []*AdditionalFee
	if i.AdditionalFees != nil {
		_ = json.Unmarshal(i.AdditionalFees, &additionalFees)
	}
	return additionalFees
}

func (i *Invoice) GetTotalPayment() uint64 {
	if i.TotalPayment != nil {
		return *i.TotalPayment
	}
	return 0
}

func (i *Invoice) GetStatus() uint32 {
	if i.Status != nil {
		return *i.Status
	}
	return 0
}

func (i *Invoice) GetCtime() uint64 {
	if i.Ctime != nil {
		return *i.Ctime
	}
	return 0
}

func GenerateInvoiceCode(merchantCode string, latestInvoice *Invoice) string {
	var (
		prefix              string
		latestInvoiceNumber int
	)

	if latestInvoice == nil {
		now := time.Now()
		date := now.Format("060102")
		prefix = merchantCode + date
	} else {
		code := latestInvoice.GetCode()
		parts := strings.Split(code, "-")
		prefix = parts[0]
		latestInvoiceNumber, _ = strconv.Atoi(parts[1])
	}

	invoiceNumber := latestInvoiceNumber + 1
	return prefix + "-" + strconv.Itoa(invoiceNumber)
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *Invoice) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
