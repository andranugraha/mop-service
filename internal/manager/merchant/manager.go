package merchant

import (
	"context"

	mDto "github.com/empnefsi/mop-service/internal/dto/merchant"
	mMod "github.com/empnefsi/mop-service/internal/module/merchant"
)

type Manager interface {
	GetMerchantByCode(ctx context.Context, code string) (*mDto.MerchantResponseData, error)
}

type impl struct {
	merchantModule mMod.Module
}

func NewManager() Manager {
	return &impl{
		merchantModule: mMod.GetModule(),
	}
}
