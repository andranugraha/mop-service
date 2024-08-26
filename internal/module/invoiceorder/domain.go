package invoiceorder

const tableName = "invoice_order_tab"

type InvoiceOrder struct {
	Id        *uint64 `gorm:"primaryKey" json:"id"`
	OrderId   *uint64 `json:"order_id"`
	InvoiceId *uint64 `json:"invoice_id"`
	Ctime     *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime     *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime     *uint64 `json:"dtime"`
}

func (i *InvoiceOrder) TableName() string {
	return tableName
}

func (i *InvoiceOrder) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *InvoiceOrder) GetOrderId() uint64 {
	if i.OrderId != nil {
		return *i.OrderId
	}
	return 0
}

func (i *InvoiceOrder) GetInvoiceId() uint64 {
	if i.InvoiceId != nil {
		return *i.InvoiceId
	}
	return 0
}
