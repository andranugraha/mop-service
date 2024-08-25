package landing

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	mManager "github.com/empnefsi/mop-service/internal/manager/merchant"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Landing(c *fiber.Ctx) error
}

type impl struct {
	manager mManager.Manager
}

func NewHandler() Handler {
	return &impl{
		manager: mManager.NewManager(),
	}
}

func (h *impl) Landing(c *fiber.Ctx) error {
	code := c.Params("code")

	if code == "" {
		return response.Error(c, constant.ErrCodeInvalidParam, "Code is required")
	}

	data, err := h.manager.GetMerchantByCode(c.UserContext(), code)
	if err != nil {
		return response.Error(c, constant.ErrCodeInternalServer, err.Error())
	}

	return response.Success(c, code, data)
}
