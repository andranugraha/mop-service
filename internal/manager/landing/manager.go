package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/module/banner"

	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/merchant"
)

type Manager interface {
	Landing(ctx context.Context, code string) (*landing.LandingResponse, error)
	GetLandingBanners(ctx context.Context, code string) (*landing.GetLandingBannersResponse, error)
}

type impl struct {
	merchantModule merchant.Module
	bannerModule   banner.Module
}

func NewManager() Manager {
	return &impl{
		merchantModule: merchant.GetModule(),
		bannerModule:   banner.GetModule(),
	}
}
