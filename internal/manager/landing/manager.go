package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/merchant"
)

type Manager interface {
	Landing(ctx context.Context, code string) (*landing.LandingResponseData, error)
}

type impl struct {
	merchantModule merchant.Module
}

func NewManager() Manager {
	return &impl{
		merchantModule: merchant.GetModule(),
	}
}
