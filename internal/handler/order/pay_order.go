package order

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/dto/order"
	"github.com/gofiber/fiber/v2"
)

func (h *impl) PayOrder(c *fiber.Ctx) error {
	req := new(order.PayOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	file, err := c.FormFile("proof_of_payment")
	if err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}
	req.ProofOfPayment = file

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
