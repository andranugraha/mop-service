package auth

import (
	"context"
	"github.com/empnefsi/mop-service/internal/dto/auth"
	"github.com/empnefsi/mop-service/internal/module/user"
)

type Manager interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponseData, error)
}

type impl struct {
	userModule user.Module
}

func NewManager() Manager {
	return &impl{
		userModule: user.GetModule(),
	}
}
