package order

import (
	"github.com/empnefsi/mop-service/internal/module/orderitem"
	"github.com/empnefsi/mop-service/internal/module/tableorder"
	"gorm.io/gorm"
	"time"
)

const tableName = "order_tab"

type Order struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	Code       *string `json:"code"`
	TotalSpend *uint64 `json:"total_spend"`
	Status     *uint32 `json:"status"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`

	Tables     []*tableorder.TableOrder `gorm:"foreignKey:OrderId;references:Id" json:"tables"`
	OrderItems []*orderitem.OrderItem   `gorm:"foreignKey:OrderId;references:Id" json:"order_items"`
}

func (i *Order) TableName() string {
	return tableName
}

func (i *Order) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *Order) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *Order) GetCode() string {
	if i.Code != nil {
		return *i.Code
	}
	return ""
}

func (i *Order) GetTotalSpend() uint64 {
	if i.TotalSpend != nil {
		return *i.TotalSpend
	}
	return 0
}

func (i *Order) GetStatus() uint32 {
	if i.Status != nil {
		return *i.Status
	}
	return 0
}

func (i *Order) GetTables() []*tableorder.TableOrder {
	if i.Tables != nil {
		return i.Tables
	}
	return nil
}

func (i *Order) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Ctime = &now
	i.Mtime = &now
	return nil
}

func (i *Order) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
