package merchant

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (i *impl) GetMerchantByCode(ctx context.Context, code string) (*Merchant, error) {
	merchant, err := i.cacheStore.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if merchant != nil {
		return merchant, nil
	}

	merchant, err = i.dbStore.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if merchant == nil {
		return nil, constant.ErrFailedToGetMerchant
	}

	err = i.cacheStore.SetMerchant(ctx, merchant)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}
