package invoice

import (
	"encoding/json"
	"time"

	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/order"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"gorm.io/gorm"
)

const tableName = "invoice_tab"

type Invoice struct {
	Id             *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId     *uint64 `json:"merchant_id"`
	PaymentTypeId  *uint64 `json:"payment_type_id"`
	Code           *string `json:"code"`
	AdditionalFees []byte  `json:"additional_fees"`
	TotalPayment   *uint64 `json:"total_payment"`
	Status         *uint32 `json:"status"`
	Ctime          *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime          *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime          *uint64 `json:"dtime"`

	PaymentType *paymenttype.PaymentType `gorm:"foreignKey:PaymentTypeId;references:Id" json:"payment_type"`
	Merchant    *merchant.Merchant       `gorm:"foreignKey:MerchantId;references:Id" json:"merchant"`
	Orders      []*order.Order           `gorm:"many2many:invoice_order_tab;" json:"orders"`
}

type AdditionalFee struct {
	Id     *uint64 `json:"id"`
	Name   *string `json:"name"`
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