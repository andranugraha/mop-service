package additionalfee

import (
	"gorm.io/gorm"
	"time"
)

const tableName = "additional_fee_tab"

const (
	AdditionalFeeTypeFixed = iota
	AdditionalFeeTypePercentage
)

type AdditionalFee struct {
	Id          *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId  *uint64 `json:"merchant_id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Type        *uint32 `json:"type"`
	Fee         *uint64 `json:"fee"`
	Ctime       *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime       *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime       *uint64 `json:"dtime"`
}

func (i *AdditionalFee) TableName() string {
	return tableName
}

func (i *AdditionalFee) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *AdditionalFee) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *AdditionalFee) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *AdditionalFee) GetDescription() string {
	if i.Description != nil {
		return *i.Description
	}
	return ""
}

func (i *AdditionalFee) GetType() uint32 {
	if i.Type != nil {
		return *i.Type
	}
	return 0
}

func (i *AdditionalFee) GetFee() uint64 {
	if i.Fee != nil {
		return *i.Fee
	}
	return 0
}

func (i *AdditionalFee) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *AdditionalFee) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
