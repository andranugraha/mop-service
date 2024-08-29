package itemvariantoption

import (
	"gorm.io/gorm"
	"time"
)

const tableName = "item_variant_option_tab"

type ItemVariantOption struct {
	Id            *uint64 `gorm:"primaryKey" json:"id"`
	ItemVariantId *uint64 `json:"item_variant_id"`
	Name          *string `json:"name"`
	Price         *uint64 `json:"price"`
	Ctime         *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime         *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime         *uint64 `json:"dtime"`
}

func (i *ItemVariantOption) TableName() string {
	return tableName
}

func (i *ItemVariantOption) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *ItemVariantOption) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *ItemVariantOption) GetPrice() uint64 {
	if i.Price != nil {
		return *i.Price
	}
	return 0
}

func (i *ItemVariantOption) GetItemVariantId() uint64 {
	if i.ItemVariantId != nil {
		return *i.ItemVariantId
	}
	return 0
}

func (i *ItemVariantOption) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *ItemVariantOption) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
