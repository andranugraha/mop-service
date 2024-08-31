package tableorder

import (
	"time"

	"gorm.io/gorm"
)

const tableName = "table_order_tab"

type TableOrder struct {
	Id      *uint64 `gorm:"primaryKey" json:"id"`
	TableId *uint64 `json:"table_id"`
	OrderId *uint64 `json:"order_id"`
	Ctime   *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime   *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime   *uint64 `json:"dtime"`
}

func (i *TableOrder) TableName() string {
	return tableName
}

func (i *TableOrder) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *TableOrder) GetTableId() uint64 {
	if i.TableId != nil {
		return *i.TableId
	}
	return 0
}

func (i *TableOrder) GetOrderId() uint64 {
	if i.OrderId != nil {
		return *i.OrderId
	}
	return 0
}

func (i *TableOrder) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *TableOrder) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
