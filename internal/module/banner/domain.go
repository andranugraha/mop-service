package banner

import (
	"time"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const tableName = "banner_tab"

const (
	VisibilityHidden = 0
	VisibilityShow   = 1
)

type Banner struct {
	Id          *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId  *uint64 `json:"merchant_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Visibility  *uint32 `json:"visibility"`
	Priority    *uint32 `json:"priority"`
	StartDate   *uint64 `json:"start_date"`
	EndDate     *uint64 `json:"end_date"`
	Ctime       *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime       *uint64 `gorm:"autoCreateTime" json:"mtime"`
	Dtime       *uint64 `json:"dtime"`
}

func (b *Banner) TableName() string {
	return tableName
}

func (b *Banner) GetId() uint64 {
	if b.Id == nil {
		return 0
	}
	return *b.Id
}

func (b *Banner) GetMerchantId() uint64 {
	if b.MerchantId == nil {
		return 0
	}
	return *b.MerchantId
}

func (b *Banner) GetTitle() string {
	if b.Title == nil {
		return ""
	}
	return *b.Title
}

func (b *Banner) GetDescription() string {
	if b.Description == nil {
		return ""
	}
	return *b.Description
}

func (b *Banner) GetImage() string {
	if b.Image == nil {
		return ""
	}
	return *b.Image
}

func (b *Banner) GetVisibility() uint32 {
	if b.Visibility == nil {
		return 0
	}
	return *b.Visibility
}

func (b *Banner) GetPriority() uint32 {
	if b.Priority == nil {
		return 0
	}
	return *b.Priority
}

func (b *Banner) GetStartDate() uint64 {
	if b.StartDate == nil {
		return 0
	}
	return *b.StartDate
}

func (b *Banner) GetEndDate() uint64 {
	if b.EndDate == nil {
		return 0
	}
	return *b.EndDate
}

func (b *Banner) getLowestActivePriority(tx *gorm.DB) *uint32 {
	var lowestActivePriority *uint32
	_ = tx.Model(&Banner{}).
		Where("merchant_id = ? AND dtime IS NULL AND end_date > ?", b.GetMerchantId(), time.Now().Unix()).
		Select("priority").
		Order("priority ASC").
		First(&lowestActivePriority)
	return lowestActivePriority
}

func (b *Banner) BeforeCreate(tx *gorm.DB) (err error) {
	lowestPriority := b.getLowestActivePriority(tx)
	if lowestPriority == nil {
		b.Priority = proto.Uint32(0)
	} else {
		b.Priority = proto.Uint32(*lowestPriority + 1)
	}

	now := uint64(time.Now().Unix())
	b.Ctime = &now
	b.Mtime = &now
	return nil
}

func (b *Banner) BeforeUpdate(tx *gorm.DB) (err error) {
	now := uint64(time.Now().Unix())
	b.Mtime = &now
	return nil
}
