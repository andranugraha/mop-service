package landing

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/gofiber/fiber/v2"
)

func (h *impl) Landing(c *fiber.Ctx) error {
	code := c.Params("code")
	req := &landing.LandingRequest{Code: code}

	if code == "" {
		return response.Error(c, req, constant.ErrCodeInvalidParam, "Code is required")
	}

	data, err := h.manager.Landing(c.UserContext(), code)
	if err != nil {
		return response.Error(c, req, constant.GetErrorCode(err), err.Error())
	}

	return response.Success(c, code, data)
}
