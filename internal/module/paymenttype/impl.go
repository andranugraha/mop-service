package paymenttype

import "context"

func (m *impl) GetActivePaymentTypesByMerchantID(ctx context.Context, merchantID uint64) ([]*PaymentType, error) {
	paymentTypes, _ := m.cacheStore.GetActivePaymentTypesByMerchantID(ctx, merchantID)
	if paymentTypes != nil {
		return paymentTypes, nil
	}

	paymentTypes, err := m.dbStore.GetActivePaymentTypesByMerchantID(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = m.cacheStore.SetPaymentTypesByMerchantID(ctx, merchantID, paymentTypes)
	}()

	return paymentTypes, nil
}
