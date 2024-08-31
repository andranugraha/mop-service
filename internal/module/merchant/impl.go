package merchant

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (i *impl) GetMerchantByID(ctx context.Context, id uint64) (*Merchant, error) {
	merchant, _ := i.cacheStore.GetMerchantByID(ctx, id)
	if merchant != nil {
		return merchant, nil
	}

	merchant, err := i.dbStore.GetMerchantByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if merchant == nil {
		return nil, constant.ErrMerchantNotFound
	}

	go func() {
		_ = i.cacheStore.SetMerchant(ctx, merchant)
	}()

	return merchant, nil
}
