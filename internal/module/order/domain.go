package order

import (
	"errors"
	"github.com/empnefsi/mop-service/internal/module/orderitem"
	"github.com/empnefsi/mop-service/internal/module/tableorder"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
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
	Code       *string `json:"code"`
	TotalSpend *uint64 `json:"total_spend"`
	Status     *uint32 `json:"status"`
	StartTime  *uint64 `json:"start_time"`
	EndTime    *uint64 `json:"end_time"`
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

func (i *Order) generateOrderCode(latestOrder *Order) string {
	var (
		prefix            string
		latestOrderNumber int
	)
	if latestOrder == nil {
		merchantCode := i.GetCode()
		now := time.Now()
		date := now.Format("060102")
		prefix = merchantCode + date
	} else {
		code := latestOrder.GetCode()
		parts := strings.Split(code, "-")
		prefix = parts[0]
		latestOrderNumber, _ = strconv.Atoi(parts[1])
	}

	orderNumber := latestOrderNumber + 1
	return prefix + "-" + strconv.Itoa(orderNumber)
}

func (i *Order) BeforeCreate(tx *gorm.DB) error {
	var todayLatestOrder *Order
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	err := tx.
		Select("id, code").
		Where("merchant_id = ?", i.GetMerchantId()).
		Where("status != ?", StatusCancelled).
		Where("dtime is null").
		Where("ctime >= ?", startOfDay.Unix()).
		Last(&todayLatestOrder).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		todayLatestOrder = nil
	}

	i.Code = proto.String(i.generateOrderCode(todayLatestOrder))

	unixNow := uint64(now.Unix())
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
