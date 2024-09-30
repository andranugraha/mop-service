package order

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/context"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
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

	userMerchantId := contextutil.GetMerchantID(ctx)
	if orderData.GetMerchantId() != userMerchantId {
		logger.Error(ctx, "pay_order", "order not found, order_id: %d", req.OrderID)
		return nil, constant.ErrOrderNotFound
	}

	if orderData.GetInvoice().GetStatus() != invoice.StatusPending {
		logger.Error(ctx, "pay_order", "invalid order status, order_id: %d, status: %s", req.OrderID, orderData.GetStatus())
		return nil, constant.ErrOrderOrInvoiceStatusInvalid
	}

	orderData.GetInvoice().Status = proto.Uint32(invoice.StatusPaid)
	orderData.Status = proto.Uint32(order.StatusPaid)

	err = m.orderModule.UpdateOrder(ctx, orderData)
	if err != nil {
		logger.Error(ctx, "pay_order", "failed to update order status, order_id: %d, error: %v", req.OrderID, err)
		return nil, err
	}

	return &dto.PayOrderResponse{
		OrderID:   orderData.GetId(),
		OrderCode: orderData.GetInvoice().GetCode(),
		Total:     orderData.GetInvoice().GetTotalPayment(),
		Status:    orderData.GetInvoice().GetStatus(),
	}, nil
}
