package order

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/logger"
	dto "github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/empnefsi/mop-service/internal/module/invoice"
	"github.com/empnefsi/mop-service/internal/module/order"
	"github.com/empnefsi/mop-service/internal/module/paymentgateway"
	"google.golang.org/protobuf/proto"
)

func (m *impl) PaymentCallback(ctx context.Context, req *dto.PaymentCallbackRequest) error {
	if req.TransactionStatus != paymentgateway.TransactionStatusSettlement {
		return nil
	}

	invoiceData, err := m.invoiceModule.GetInvoiceByCode(ctx, req.OrderID)
	if err != nil {
		return nil
	}

	if invoiceData.GetStatus() != invoice.StatusPending {
		return nil
	}

	orderData, err := m.orderModule.GetOrderByInvoiceID(ctx, invoiceData.GetId())
	if err != nil {
		return nil
	}

	orderData.Status = proto.Uint32(order.StatusPaid)
	invoiceData.Status = proto.Uint32(invoice.StatusPaid)
	orderData.Invoice = invoiceData

	err = m.orderModule.UpdateOrder(ctx, orderData)
	if err != nil {
		logger.Error(ctx, "payment_callback", "failed to update order: %v", err.Error())
		return err
	}

	go m.PushPaymentEvent(ctx, orderData.GetId(), "Payment has been received")
	return nil
}

func (m *impl) PushPaymentEvent(ctx context.Context, orderId uint64, message string) {
	m.clientsMutex.Lock()
	defer m.clientsMutex.Unlock()
	if ch, ok := m.clients[orderId]; ok {
		logger.Info(ctx, "push_payment_event", "pushing message to order id %d: %s", orderId, message)
		ch <- message
	}
}

func (m *impl) RegisterPaymentEvent(ctx context.Context, orderId uint64) chan string {
	m.clientsMutex.Lock()
	defer m.clientsMutex.Unlock()

	logger.Info(ctx, "register_payment_event", "registering order id %d", orderId)
	ch := make(chan string)
	m.clients[orderId] = ch
	return ch
}

func (m *impl) UnregisterPaymentEvent(ctx context.Context, orderId uint64) {
	m.clientsMutex.Lock()
	defer m.clientsMutex.Unlock()

	logger.Info(ctx, "unregister_payment_event", "unregistering order id %d", orderId)
	delete(m.clients, orderId)
}
