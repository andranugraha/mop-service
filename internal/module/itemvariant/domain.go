package itemvariant

import "github.com/empnefsi/mop-service/internal/module/itemvariantoption"

const tableName = "item_variant_tab"

type ItemVariant struct {
	Id          *uint64 `gorm:"primaryKey" json:"id"`
	ItemId      *uint64 `json:"item_id"`
	Name        *string `json:"name"`
	IsRequired  *uint   `json:"is_required"`
	SelectCount *uint32 `json:"select_count"`
	SelectType  *uint   `json:"select_type"`
	Ctime       *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime       *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime       *uint64 `json:"dtime"`

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

func (i *ItemVariant) GetIsRequired() uint {
	if i.IsRequired != nil {
		return *i.IsRequired
	}
	return 0
}

func (i *ItemVariant) GetSelectCount() uint32 {
	if i.SelectCount != nil {
		return *i.SelectCount
	}
	return 0
}

func (i *ItemVariant) GetSelectType() uint {
	if i.SelectType != nil {
		return *i.SelectType
	}
	return 0
}

func (i *ItemVariant) GetOptions() []*itemvariantoption.ItemVariantOption {
	return i.Options
}
