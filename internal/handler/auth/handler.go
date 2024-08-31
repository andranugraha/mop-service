package auth

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/common/validator"
	"github.com/empnefsi/mop-service/internal/dto/auth"
	authManager "github.com/empnefsi/mop-service/internal/manager/auth"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Login(c *fiber.Ctx) error
}

type impl struct {
	manager authManager.Manager
}

func NewHandler() Handler {
	return &impl{
		manager: authManager.NewManager(),
	}
}

func (h *impl) Login(c *fiber.Ctx) error {
	req := new(auth.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	if err := validator.Validate(req); err != nil {
		return response.Error(c, req, constant.ErrCodeInvalidParam, err.Error())
	}

	data, err := h.manager.Login(c.UserContext(), req)
	if err != nil {
		code := constant.GetErrorCode(err)
		if code != constant.ErrCodeInternalServer {
			return response.Error(c, req, code, err.Error())
		}
		return response.Error(c, req, code, constant.ErrInternalServer.Error())
	}

	return response.Success(c, req, data)
}
