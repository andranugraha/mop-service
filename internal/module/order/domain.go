package order

import (
	"time"

	"github.com/empnefsi/mop-service/internal/module/invoice"

	"github.com/empnefsi/mop-service/internal/module/orderitem"
	"github.com/empnefsi/mop-service/internal/module/tableorder"
	"gorm.io/gorm"
)

const tableName = "order_tab"

const (
	StatusPending uint32 = iota
	StatusPaid
	StatusOnProcess
	StatusDone
	StatusCancelled
)

type Order struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	InvoiceId  *uint64 `json:"invoice_id"`
	TotalSpend *uint64 `json:"total_spend"`
	Status     *uint32 `json:"status"`
	StartTime  *uint64 `json:"start_time"`
	EndTime    *uint64 `json:"end_time"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`

	Tables     []*tableorder.TableOrder `gorm:"foreignKey:OrderId;references:Id" json:"tables"`
	OrderItems []*orderitem.OrderItem   `gorm:"foreignKey:OrderId;references:Id" json:"order_items"`
	Invoice    *invoice.Invoice         `gorm:"foreignKey:InvoiceId;references:Id" json:"invoice"`
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

func (i *Order) GetStartTime() uint64 {
	if i.StartTime != nil {
		return *i.StartTime
	}
	return 0
}

func (i *Order) GetEndTime() uint64 {
	if i.EndTime != nil {
		return *i.EndTime
	}
	return 0
}

func (i *Order) GetCtime() uint64 {
	if i.Ctime != nil {
		return *i.Ctime
	}
	return 0
}

func (i *Order) GetMtime() uint64 {
	if i.Mtime != nil {
		return *i.Mtime
	}
	return 0
}

func (i *Order) GetTables() []*tableorder.TableOrder {
	if i.Tables != nil {
		return i.Tables
	}
	return nil
}

func (i *Order) GetOrderItems() []*orderitem.OrderItem {
	if i.OrderItems != nil {
		return i.OrderItems
	}
	return nil
}

func (i *Order) GetInvoice() *invoice.Invoice {
	if i.Invoice != nil {
		return i.Invoice
	}
	return nil
}

func (i *Order) BeforeCreate(tx *gorm.DB) error {
	unixNow := uint64(time.Now().Unix())
	i.StartTime = &unixNow
	i.Ctime = &unixNow
	i.Mtime = &unixNow
	return nil
}

func (i *Order) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	i.Mtime = &now
	return nil
}
