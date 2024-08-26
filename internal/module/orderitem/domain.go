package orderitem

import (
	"encoding/json"
	"github.com/empnefsi/mop-service/internal/module/order"
)

const tableName = "order_item_tab"

type OrderItem struct {
	Id           *uint64 `gorm:"primaryKey" json:"id"`
	OrderId      *uint64 `json:"order_id"`
	ItemId       *uint64 `json:"item_id"`
	Name         *string `json:"name"`
	Amount       *uint64 `json:"amount"`
	ItemOptions  []byte  `json:"item_options"`
	Note         *string `json:"note"`
	PricePerItem *uint64 `json:"price_per_item"`
	TotalPrice   *uint64 `json:"total_price"`
	Ctime        *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime        *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime        *uint64 `json:"dtime"`

	Order *order.Order `gorm:"foreignKey:OrderId;references:Id" json:"order"`
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

func (i *OrderItem) GetName() string {
	if i.Name != nil {
		return *i.Name
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
