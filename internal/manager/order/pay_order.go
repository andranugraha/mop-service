package order

import (
	"context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/common/upload"
	dto "github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/empnefsi/mop-service/internal/module/invoice"
	"github.com/empnefsi/mop-service/internal/module/order"
	"google.golang.org/protobuf/proto"
)

func (m *impl) PayOrder(ctx context.Context, req *dto.PayOrderRequest) (*dto.PayOrderResponse, error) {
	orderData, err := m.orderModule.GetFullOrderDataByID(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	if orderData.GetInvoice().GetStatus() != invoice.StatusPending {
		logger.Error(ctx, "pay_order", "invalid order status, order_id: %d, status: %s", req.OrderID, orderData.GetStatus())
		return nil, constant.ErrOrderOrInvoiceStatusInvalid
	}

	merchantData, err := m.merchantModule.GetMerchantByID(ctx, orderData.GetMerchantId())
	if err != nil {
		return nil, err
	}

	fileName := "invoice_" + orderData.GetInvoice().GetCode()
	uploadedFilePath, err := upload.File(req.ProofOfPayment, fileName, merchantData.GetCode())
	if err != nil {
		logger.Error(ctx, "pay_order", "failed to upload proof of payment, order_id: %d, error: %v", req.OrderID, err)
		return nil, err
	}

	orderData.GetInvoice().PaymentProof = uploadedFilePath
	orderData.GetInvoice().Status = proto.Uint32(invoice.StatusPaid)
	orderData.Status = proto.Uint32(order.StatusPaid)

	err = m.orderModule.UpdateOrder(ctx, orderData)
	if err != nil {
		logger.Error(ctx, "pay_order", "failed to update order status, order_id: %d, error: %v", req.OrderID, err)
		return nil, err
	}

	return &dto.PayOrderResponse{
		InvoiceID:   orderData.GetInvoice().GetId(),
		InvoiceCode: orderData.GetInvoice().GetCode(),
	}, nil
}
