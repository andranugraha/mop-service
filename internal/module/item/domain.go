package item

import (
	"time"

	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"gorm.io/gorm"
)

const tableName = "item_tab"

type Item struct {
	Id             *uint64 `gorm:"primaryKey" json:"id"`
	ItemCategoryId *uint64 `json:"item_category_id"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	Price          *uint64 `json:"price"`
	Ctime          *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime          *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime          *uint64 `json:"dtime"`

	Variants []*itemvariant.ItemVariant `gorm:"foreignKey:ItemId;references:Id" json:"variants"`
}

func (i *Item) TableName() string {
	return tableName
}

func (i *Item) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *Item) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *Item) GetDescription() string {
	if i.Description != nil {
		return *i.Description
	}
	return ""
}

func (i *Item) GetPrice() uint64 {
	if i.Price != nil {
		return *i.Price
	}
	return 0
}

func (i *Item) GetItemCategoryId() uint64 {
	if i.ItemCategoryId != nil {
		return *i.ItemCategoryId
	}
	return 0
}

func (i *Item) GetVariants() []*itemvariant.ItemVariant {
	return i.Variants
}

func (i *Item) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *Item) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
