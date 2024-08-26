package landing

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/dto/landing"
)

func (m *impl) Landing(ctx context.Context, code string) (*landing.LandingResponseData, error) {
	merchantData, err := m.merchantModule.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if merchantData == nil {
		return nil, constant.ErrFailedToGetMerchant
	}

	return &landing.LandingResponseData{
		Code: merchantData.GetCode(),
		Name: merchantData.GetName(),
	}, nil
}
