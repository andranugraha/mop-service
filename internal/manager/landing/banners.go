package landing

import (
	"context"
	"time"

	"github.com/empnefsi/mop-service/internal/config"
	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/empnefsi/mop-service/internal/module/banner"
)

func (m *impl) GetLandingBanners(ctx context.Context, code string) (*landing.GetLandingBannersResponse, error) {
	merchantData, err := m.merchantModule.GetMerchantByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	banners, err := m.bannerModule.GetActiveBannersByMerchantID(ctx, merchantData.GetId())
	if err != nil {
		return nil, err
	}

	return &landing.GetLandingBannersResponse{
		Banners: m.filterAndConvertBanners(banners),
	}, nil
}

func (m *impl) filterAndConvertBanners(banners []*banner.Banner) []landing.Banner {
	var filteredBanners []landing.Banner
	now := time.Now().Unix()
	for _, bannerData := range banners {
		if bannerData.GetVisibility() == banner.VisibilityHidden || bannerData.GetStartDate() > uint64(now) {
			continue
		}
		filteredBanners = append(filteredBanners, landing.Banner{
			Id:          bannerData.GetId(),
			Title:       bannerData.GetTitle(),
			Description: bannerData.GetDescription(),
			Image:       config.GetCDNEndpoint() + "/" + bannerData.GetImage(),
			StartDate:   bannerData.GetStartDate(),
			EndDate:     bannerData.EndDate,
		})
	}
	return filteredBanners
}
