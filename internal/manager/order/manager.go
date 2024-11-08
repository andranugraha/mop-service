package order

import (
	"context"
	"sync"

	"github.com/empnefsi/mop-service/internal/module/invoice"

	dto "github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/empnefsi/mop-service/internal/module/item"
	"github.com/empnefsi/mop-service/internal/module/itemvariant"
	"github.com/empnefsi/mop-service/internal/module/itemvariantoption"
	"github.com/empnefsi/mop-service/internal/module/merchant"
	"github.com/empnefsi/mop-service/internal/module/order"
	"github.com/empnefsi/mop-service/internal/module/table"
)

type Manager interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error)
	PayOrder(ctx context.Context, req *dto.PayOrderRequest) (*dto.PayOrderResponse, error)
	PaymentCallback(ctx context.Context, req *dto.PaymentCallbackRequest) error
	PushPaymentEvent(ctx context.Context, orderId uint64, message string)
	RegisterPaymentEvent(ctx context.Context, orderId uint64) chan string
	UnregisterPaymentEvent(ctx context.Context, orderId uint64)
}

type impl struct {
	merchantModule          merchant.Module
	orderModule             order.Module
	tableModule             table.Module
	itemModule              item.Module
	itemVariantModule       itemvariant.Module
	itemVariantOptionModule itemvariantoption.Module
	invoiceModule           invoice.Module

	clients      map[uint64]chan string
	clientsMutex sync.Mutex
}

func NewManager() Manager {
	manager := &impl{
		merchantModule:          merchant.GetModule(),
		orderModule:             order.GetModule(),
		tableModule:             table.GetModule(),
		itemModule:              item.GetModule(),
		itemVariantModule:       itemvariant.GetModule(),
		itemVariantOptionModule: itemvariantoption.GetModule(),
		invoiceModule:           invoice.GetModule(),
		clients:                 make(map[uint64]chan string),
	}
	return manager
}
