package orderitem

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

const tableName = "order_item_tab"

type OrderItem struct {
	Id           *uint64 `gorm:"primaryKey" json:"id"`
	OrderId      *uint64 `json:"order_id"`
	ItemId       *uint64 `json:"item_id"`
	ItemName     *string `json:"item_name"`
	Amount       *uint64 `json:"amount"`
	ItemOptions  []byte  `json:"item_options"`
	Note         *string `json:"note"`
	PricePerItem *uint64 `json:"price_per_item"`
	TotalPrice   *uint64 `json:"total_price"`
	Ctime        *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime        *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime        *uint64 `json:"dtime"`
}

type Variant struct {
	Id          *uint64   `json:"variant_id"`
	VariantName *string   `json:"variant_name"`
	Options     []Options `json:"options"`
}

type Options struct {
	Id         *uint64 `json:"option_id"`
	OptionName *string `json:"option_name"`
	Price      *uint64 `json:"price"`
}

func (i *OrderItem) TableName() string {
	return tableName
}

func (i *OrderItem) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *OrderItem) GetOrderId() uint64 {
	if i.OrderId != nil {
		return *i.OrderId
	}
	return 0
}

func (i *OrderItem) GetItemId() uint64 {
	if i.ItemId != nil {
		return *i.ItemId
	}
	return 0
}

func (i *OrderItem) GetItemName() string {
	if i.ItemName != nil {
		return *i.ItemName
	}
	return ""
}

func (i *OrderItem) GetAmount() uint64 {
	if i.Amount != nil {
		return *i.Amount
	}
	return 0
}

func (i *OrderItem) GetItemOptions() []*Variant {
	var variants []*Variant
	if i.ItemOptions != nil {
		_ = json.Unmarshal(i.ItemOptions, &variants)
	}
	return variants
}

func (i *OrderItem) GetNote() string {
	if i.Note != nil {
		return *i.Note
	}
	return ""
}

func (i *OrderItem) GetPricePerItem() uint64 {
	if i.PricePerItem != nil {
		return *i.PricePerItem
	}
	return 0
}

func (i *OrderItem) GetTotalPrice() uint64 {
	if i.TotalPrice != nil {
		return *i.TotalPrice
	}
	return 0
}

func (i *OrderItem) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *OrderItem) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
