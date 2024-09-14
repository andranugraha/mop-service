package merchant

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/logger"
	dto "github.com/empnefsi/mop-service/internal/dto/merchant"
	"github.com/empnefsi/mop-service/internal/module/additionalfee"
	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/paymenttype"
)

type Manager interface {
	GetMerchantActivePaymentTypes(ctx context.Context, merchantID uint64) (*dto.GetMerchantActivePaymentTypesResponse, error)
	GetMerchantActiveAdditionalFees(ctx context.Context, merchantID uint64) (*dto.GetMerchantActiveAdditionalFeesResponse, error)
}

type impl struct {
	merchantModule      merchant.Module
	paymentTypeModule   paymenttype.Module
	additionalFeeModule additionalfee.Module
}

func NewManager() Manager {
	manager := &impl{
		merchantModule:      merchant.GetModule(),
		paymentTypeModule:   paymenttype.GetModule(),
		additionalFeeModule: additionalfee.GetModule(),
	}
	return manager
}

func (m *impl) GetMerchantActivePaymentTypes(ctx context.Context, merchantID uint64) (*dto.GetMerchantActivePaymentTypesResponse, error) {
	paymentTypes, err := m.paymentTypeModule.GetActivePaymentTypesByMerchantID(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	if len(paymentTypes) == 0 {
		logger.Error(ctx, "get_merchant_active_payment_types", "no payment types found for merchant %d", merchantID)
		return nil, nil
	}

	activePaymentTypes := make([]*dto.PaymentType, 0)
	for _, paymentType := range paymentTypes {
		activePaymentTypes = append(activePaymentTypes, &dto.PaymentType{
			Type: paymentType.GetType(),
			Name: paymentType.GetName(),
		})
	}

	return &dto.GetMerchantActivePaymentTypesResponse{
		PaymentTypes: activePaymentTypes,
	}, nil
}

func (m *impl) GetMerchantActiveAdditionalFees(ctx context.Context, merchantID uint64) (*dto.GetMerchantActiveAdditionalFeesResponse, error) {
	additionalFees, err := m.additionalFeeModule.GetActiveAdditionalFeesByMerchantID(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	if len(additionalFees) == 0 {
		logger.Error(ctx, "get_merchant_active_additional_fees", "no additional fees found for merchant %d", merchantID)
		return nil, nil
	}

	activeAdditionalFees := make([]*dto.AdditionalFee, 0)
	for _, additionalFee := range additionalFees {
		activeAdditionalFees = append(activeAdditionalFees, &dto.AdditionalFee{
			Name:        additionalFee.GetName(),
			Description: additionalFee.GetDescription(),
			Type:        additionalFee.GetType(),
			Fee:         additionalFee.GetFee(),
		})
	}

	return &dto.GetMerchantActiveAdditionalFeesResponse{
		AdditionalFees: activeAdditionalFees,
	}, nil
}
