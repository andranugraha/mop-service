package merchant

import (
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

const tableName = "merchant_tab"

type Merchant struct {
	Id    *uint64 `gorm:"primaryKey" json:"id"`
	Code  *string `gorm:"unique" json:"code"`
	Name  *string `json:"name"`
	Ctime *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime *uint64 `json:"dtime"`
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

// BeforeCreate is a GORM hook that runs before a record is created in the database
func (m *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	for {
		m.Code = proto.String(generateMerchantCode(m.GetName()))
		var count int64
		tx.Model(&Merchant{}).Where("code = ?", m.Code).Count(&count)
		if count == 0 {
			break
		}
	}
	return nil
}
