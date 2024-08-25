package merchant

import (
	"context"
)

func (i *impl) GetMerchantByCode(ctx context.Context, code string) (*Merchant, error) {
	merchant, err := i.dbStore.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return merchant, nil
}
