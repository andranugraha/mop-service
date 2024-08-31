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
		return nil, constant.ErrItemNotFound
	}

	err = i.cacheStore.SetMerchant(ctx, merchant)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}

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
