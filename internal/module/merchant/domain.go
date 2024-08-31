package merchant

import (
	"math/rand"
	"strings"
	"time"

	"github.com/empnefsi/mop-service/internal/module/additionalfee"
	"github.com/empnefsi/mop-service/internal/module/itemcategory"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
	"github.com/empnefsi/mop-service/internal/module/table"
	"github.com/empnefsi/mop-service/internal/module/user"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const tableName = "merchant_tab"

type Merchant struct {
	Id    *uint64 `gorm:"primaryKey" json:"id"`
	Code  *string `gorm:"unique" json:"code"`
	Name  *string `json:"name"`
	Ctime *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime *uint64 `json:"dtime"`

	Users          []*user.User                   `gorm:"foreignKey:MerchantId;references:Id" json:"users"`
	PaymentTypes   []*paymenttype.PaymentType     `gorm:"foreignKey:MerchantId;references:Id" json:"payment_types"`
	AdditionalFees []*additionalfee.AdditionalFee `gorm:"foreignKey:MerchantId;references:Id" json:"additional_fees"`
	Tables         []*table.Table                 `gorm:"foreignKey:MerchantId;references:Id" json:"tables"`
	ItemCategories []*itemcategory.ItemCategory   `gorm:"foreignKey:MerchantId;references:Id" json:"item_categories"`
}

func (m *Merchant) TableName() string {
	return tableName
}

func (m *Merchant) GetId() uint64 {
	if m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Merchant) GetCode() string {
	if m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *Merchant) GetName() string {
	if m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Merchant) GetCtime() uint64 {
	if m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

func (m *Merchant) GetMtime() uint64 {
	if m.Mtime != nil {
		return *m.Mtime
	}
	return 0
}

func (m *Merchant) GetAdditionalFees() []*additionalfee.AdditionalFee {
	if m.AdditionalFees != nil {
		return m.AdditionalFees
	}
	return nil
}

func (m *Merchant) GetPaymentTypes() []*paymenttype.PaymentType {
	if m.PaymentTypes != nil {
		return m.PaymentTypes
	}
	return nil
}

func generateMerchantCode(name string) string {
	words := strings.Fields(name)
	var initials string
	for _, word := range words {
		initials += strings.ToUpper(string(word[0]))
	}

	lengthNeeded := 8 - len(initials)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomChars := make([]byte, lengthNeeded)
	for i := range randomChars {
		randomChars[i] = charset[rand.Intn(len(charset))]
	}

	return initials + string(randomChars)
}

func (m *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	for {
		m.Code = proto.String(generateMerchantCode(m.GetName()))
		var count int64
		tx.Model(&Merchant{}).Where("code = ?", m.Code).Count(&count)
		if count == 0 {
			break
		}
	}

	now := uint64(time.Now().Unix())
	m.Ctime = &now
	m.Mtime = &now
	return nil
}

func (m *Merchant) BeforeUpdate(tx *gorm.DB) (err error) {
	now := uint64(time.Now().Unix())
	m.Mtime = &now
	return nil
}
