package user

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (i *impl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user, err := i.cacheStore.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	user, err = i.dbStore.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, constant.ErrInvalidIdentifierOrPassword
	}

	err = i.cacheStore.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
