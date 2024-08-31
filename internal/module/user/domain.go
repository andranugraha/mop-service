package user

import (
	"gorm.io/gorm"
	"time"
)

const tableName = "user_tab"

type User struct {
	Id         *uint64 `gorm:"primaryKey" json:"id"`
	MerchantId *uint64 `json:"merchant_id"`
	Email      *string `gorm:"uniqueKey" json:"email"`
	Password   *string `json:"password"`
	Ctime      *uint64 `gorm:"autoCreateTime" json:"ctime"`
	Mtime      *uint64 `gorm:"autoUpdateTime" json:"mtime"`
	Dtime      *uint64 `json:"dtime"`
}

func (u *User) TableName() string {
	return tableName
}

func (u *User) GetId() uint64 {
	if u.Id != nil {
		return *u.Id
	}
	return 0
}

func (u *User) GetEmail() string {
	if u.Email != nil {
		return *u.Email
	}
	return ""
}

func (u *User) GetPassword() string {
	if u.Password != nil {
		return *u.Password
	}
	return ""
}

func (u *User) GetMerchantId() uint64 {
	if u.MerchantId != nil {
		return *u.MerchantId
	}
	return 0
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	u.Ctime = &now
	u.Mtime = &now
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	now := uint64(time.Now().Unix())
	u.Mtime = &now
	return nil
}
