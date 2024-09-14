package additionalfee

import "context"

func (m *impl) GetActiveAdditionalFeesByMerchantID(ctx context.Context, merchantID uint64) ([]*AdditionalFee, error) {
	additionalFees, _ := m.cacheStore.GetActiveAdditionalFeesByMerchantID(ctx, merchantID)
	if additionalFees != nil {
		return additionalFees, nil
	}

	additionalFees, err := m.dbStore.GetActiveAdditionalFeesByMerchantID(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = m.cacheStore.SetAdditionalFeeByMerchantID(ctx, merchantID, additionalFees)
	}()

	return additionalFees, nil
}
