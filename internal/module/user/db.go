package user

import (
	"context"
	"errors"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"gorm.io/gorm"
)

type db struct {
	client *gorm.DB
}

func (d *db) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := d.client.
		Select("id, email, password, merchant_id").
		Where("email = ?", email).
		Where("dtime is null").
		Take(&user).
		Error
	if err != nil {
		logger.Error(ctx, "fetch_user_from_db", "failed to fetch user: %v", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
