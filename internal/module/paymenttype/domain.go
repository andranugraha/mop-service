package paymenttype

const tableName = "payment_type_tab"

type PaymentType struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	Type       *string `json:"type"`
	Asset      *string `json:"asset"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`
}

func (i *PaymentType) TableName() string {
	return tableName
}

func (i *PaymentType) GetId() uint64 {
	if i.Id != nil {
		return *i.Id
	}
	return 0
}

func (i *PaymentType) GetMerchantId() uint64 {
	if i.MerchantId != nil {
		return *i.MerchantId
	}
	return 0
}

func (i *PaymentType) GetType() string {
	if i.Type != nil {
		return *i.Type
	}
	return ""
}

func (i *PaymentType) GetAsset() string {
	if i.Asset != nil {
		return *i.Asset
	}
	return ""
}
