package itemcategory

import (
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}