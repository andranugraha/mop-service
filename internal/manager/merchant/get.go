package merchant

import (
	"context"

	mDto "github.com/empnefsi/mop-service/internal/dto/merchant"
)

func (i *impl) GetMerchantByCode(ctx context.Context, code string) (*mDto.MerchantResponseData, error) {
	merchant, err := i.merchantModule.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &mDto.MerchantResponseData{
		Code: *merchant.Code,
		Name: *merchant.Name,
	}, nil
}
