package paymenttype

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

const tableName = "payment_type_tab"

const (
	PaymentTypeCashier = iota
	PaymentTypeQR
)

type PaymentType struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	Type       *uint32 `json:"type"`
	Name       *string `json:"name"`
	ExtraData  []byte  `json:"extra_data"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`
}

type QRPaymentTypeExtraData struct {
	ImageURL string `json:"image_url"`
}

func (i *PaymentType) TableName() string {
	return tableName
}

func (i *PaymentType) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *PaymentType) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *PaymentType) GetType() uint32 {
	if i.Type != nil {
		return *i.Type
	}
	return 0
}

func (i *PaymentType) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *PaymentType) GetQRPaymentTypeExtraData() *QRPaymentTypeExtraData {
	if i.GetType() == PaymentTypeQR {
		extraData := &QRPaymentTypeExtraData{}
		_ = json.Unmarshal(i.ExtraData, extraData)
		return extraData
	}
	return nil
}

func (i *PaymentType) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *PaymentType) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
