package merchant

import (
	"strconv"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/manager/merchant"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetMerchantActivePaymentTypes(c *fiber.Ctx) error
	GetMerchantActiveAdditionalFees(c *fiber.Ctx) error
}

type impl struct {
	manager merchant.Manager
}

func NewHandler() Handler {
	return &impl{
		manager: merchant.NewManager(),
	}
}

func (h *impl) GetMerchantActivePaymentTypes(c *fiber.Ctx) error {
	merchantId := c.Params("merchant_id")
	if merchantId == "" {
		return response.Error(c, nil, constant.ErrCodeInvalidParam, "merchant_id is required")
	}

	convMerchantId, _ := strconv.Atoi(merchantId)
	req := map[string]interface{}{
		"merchant_id": convMerchantId,
	}
	data, err := h.manager.GetMerchantActivePaymentTypes(c.UserContext(), uint64(convMerchantId))
	if err != nil {
		code := constant.GetErrorCode(err)

		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, data)
}

func (h *impl) GetMerchantActiveAdditionalFees(c *fiber.Ctx) error {
	merchantId := c.Params("merchant_id")
	if merchantId == "" {
		return response.Error(c, nil, constant.ErrCodeInvalidParam, "merchant_id is required")
	}

	convMerchantId, _ := strconv.Atoi(merchantId)
	req := map[string]interface{}{
		"merchant_id": convMerchantId,
	}
	data, err := h.manager.GetMerchantActiveAdditionalFees(c.UserContext(), uint64(convMerchantId))
	if err != nil {
		code := constant.GetErrorCode(err)

		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, data)
}
