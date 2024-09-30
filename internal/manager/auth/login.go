package auth

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/strings"
	"github.com/empnefsi/mop-service/internal/dto/auth"
	"github.com/empnefsi/mop-service/internal/module/user"
)

func (i *impl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	userData, err := i.userModule.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if userData == nil {
		return nil, constant.ErrInvalidIdentifierOrPassword
	}

	if !i.validPassword(req.Password, userData.GetPassword()) {
		return nil, constant.ErrInvalidIdentifierOrPassword
	}

	token, err := i.generateToken(userData)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Token: token,
	}, nil
}

func (i *impl) validPassword(reqPassword, hashedPassword string) bool {
	return strings.CheckPasswordHash(reqPassword, hashedPassword)
}

func (i *impl) generateToken(payload *user.User) (string, error) {
	return strings.GenerateToken(strings.Claims{
		UserID:     payload.GetId(),
		Email:      payload.GetEmail(),
		MerchantID: payload.GetMerchantId(),
		Role:       payload.GetRole(),
	})
}
