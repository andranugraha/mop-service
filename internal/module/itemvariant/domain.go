package itemvariant

import (
	"time"

	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"gorm.io/gorm"
)

const tableName = "item_variant_tab"

type ItemVariant struct {
	Id        *uint64 `gorm:"primaryKey" json:"id"`
	ItemId    *uint64 `json:"item_id"`
	Name      *string `json:"name"`
	MinSelect *uint32 `json:"min_select"`
	MaxSelect *uint32 `json:"max_select"`
	Ctime     *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime     *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime     *uint64 `json:"dtime"`

	Options []*itemvariantoption.ItemVariantOption `gorm:"foreignKey:ItemVariantId;references:Id" json:"options"`
}

func (i *ItemVariant) TableName() string {
	return tableName
}

func (i *ItemVariant) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *ItemVariant) GetName() string {
	if i.Name != nil {
		return *i.Name
	}
	return ""
}

func (i *ItemVariant) GetItemId() uint64 {
	if i.ItemId != nil {
		return *i.ItemId
	}
	return 0
}

func (i *ItemVariant) GetMinSelect() uint32 {
	if i.MinSelect != nil {
		return *i.MinSelect
	}
	return 0
}

func (i *ItemVariant) GetMaxSelect() uint32 {
	if i.MaxSelect != nil {
		return *i.MaxSelect
	}
	return 0
}

func (i *ItemVariant) GetOptions() []*itemvariantoption.ItemVariantOption {
	return i.Options
}

func (i *ItemVariant) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *ItemVariant) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
