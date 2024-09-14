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

	return nil
}
