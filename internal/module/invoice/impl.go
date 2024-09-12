package invoice

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/constant"
	"gorm.io/gorm"
)

func (m *impl) GetInvoiceByID(ctx context.Context, id uint64) (*Invoice, error) {
	invoiceData, err := m.dbStore.GetInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if invoiceData == nil {
		return nil, constant.ErrOrderInvoiceNotFound
	}

	return invoiceData, nil
}

func (m *impl) UpdateInvoiceTx(ctx context.Context, tx *gorm.DB, invoice *Invoice) error {
	return m.dbStore.UpdateInvoiceTx(ctx, tx, invoice)
}

func (m *impl) GetTodayLatestInvoice(ctx context.Context, merchantID uint64) (*Invoice, error) {
	invoiceData, err := m.dbStore.GetTodayLatestInvoice(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	if invoiceData == nil {
		return nil, constant.ErrOrderInvoiceNotFound
	}

	return invoiceData, nil
}
