package order

import "context"

func (m *impl) CreateOrder(ctx context.Context, order *Order) error {
	return m.dbStore.CreateOrder(ctx, order)
}
