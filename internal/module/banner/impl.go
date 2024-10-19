package banner

import (
	"context"
)

func (i *impl) GetActiveBannersByMerchantID(ctx context.Context, merchantID uint64) ([]*Banner, error) {
	banners, _ := i.cacheStore.GetActiveBannersByMerchantID(ctx, merchantID)
	if banners != nil {
		return banners, nil
	}

	banners, err := i.dbStore.GetActiveBannersByMerchantID(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	if banners == nil || len(banners) == 0 {
		return nil, nil
	}

	go func() {
		_ = i.cacheStore.SetActiveBannersByMerchantID(ctx, merchantID, banners)
	}()

	return banners, nil
}
