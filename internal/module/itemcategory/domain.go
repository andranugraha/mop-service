package itemcategory

import (
	"time"

	"github.com/empnefsi/mop-service/internal/module/item"
	"gorm.io/gorm"
)

const tableName = "item_category_tab"

type ItemCategory struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	Name       *string `json:"name"`
	Priority   *int32  `json:"priority"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`

	Items []*item.Item `gorm:"foreignKey:ItemCategoryId;references:Id" json:"items"`
}

func (i *ItemCategory) TableName() string {
	return tableName
}

func (i *ItemCategory) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *ItemCategory) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *ItemCategory) GetPriority() int32 {
	if i.Priority != nil {
		return *i.Priority
	}
	return 0
}

func (i *ItemCategory) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *ItemCategory) GetItems() []*item.Item {
	return i.Items
}

func (i *ItemCategory) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *ItemCategory) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
