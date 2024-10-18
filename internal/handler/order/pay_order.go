package order

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/common/validator"
	"github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/gofiber/fiber/v2"
)

// @Summary Pay Order
// @Description Pay order
// @Tags Order
// @Accept json
// @Produce json
// @Param body body PayOrderRequest true "Pay order request"
// @Success 200 {object} PayOrderResponse
// @Router /api/v1/order/pay [post]
func (h *impl) PayOrder(c *fiber.Ctx) error {
	req := new(order.PayOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	if err := validator.Validate(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	data, err := h.manager.PayOrder(c.UserContext(), req)
	if err != nil {
		code := constant.GetErrorCode(err)
		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, data)
}
