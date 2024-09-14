package order

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
)

func (m *impl) CreateOrder(ctx context.Context, order *Order) error {
	return m.dbStore.CreateOrder(ctx, order)
}

func (m *impl) GetFullOrderDataByID(ctx context.Context, id uint64) (*Order, error) {
	orderData, err := m.dbStore.GetFullOrderDataByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if orderData == nil {
		return nil, constant.ErrOrderNotFound
	}

	return orderData, nil
}

func (m *impl) UpdateOrder(ctx context.Context, order *Order) error {
	return m.dbStore.UpdateOrder(ctx, order)
}

func (m *impl) GetOrderByInvoiceID(ctx context.Context, invoiceId uint64) (*Order, error) {
	orderData, err := m.dbStore.GetOrderByInvoiceID(ctx, invoiceId)
	if err != nil {
		return nil, err
	}

	if orderData == nil {
		return nil, constant.ErrOrderNotFound
	}

	return orderData, nil
}
