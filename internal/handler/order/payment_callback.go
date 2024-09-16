package order

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *impl) PaymentCallback(c *fiber.Ctx) error {
	req := new(order.PaymentCallbackRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	err := h.manager.PaymentCallback(c.UserContext(), req)
	if err != nil {
		code := constant.GetErrorCode(err)
		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, nil)
}

func (h *impl) PushPaymentEvent(c *fiber.Ctx) error {
	orderId := c.Params("order_id")
	if orderId == "" {
		return response.Error(c, nil, constant.ErrCodeInvalidParam, "order_id is required")
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	uint64OrderId, _ := strconv.ParseUint(orderId, 10, 64)
	ch := h.manager.RegisterPaymentEvent(c.UserContext(), uint64OrderId)
	defer func() {
		h.manager.UnregisterPaymentEvent(c.UserContext(), uint64OrderId)
		close(ch)
	}()

	for msg := range ch {
		c.Response().BodyWriter().Write([]byte("data: " + msg + "\n\n"))
	}

	return nil
}
