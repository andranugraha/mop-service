package order

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/common/validator"
	"github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/gofiber/fiber/v2"
)

func (h *impl) CreateOrder(c *fiber.Ctx) error {
	req := new(order.CreateOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	if err := validator.Validate(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	data, err := h.manager.CreateOrder(c.UserContext(), req)
	if err != nil {
		code := constant.GetErrorCode(err)
		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, data)
}
