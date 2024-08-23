package auth

import (
	"errors"
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
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
		return response.Error(c, constant.ErrCodeInvalidParam, err.Error())
	}

	if err := validator.Validate(req); err != nil {
		return response.Error(c, constant.ErrCodeInvalidParam, err.Error())
	}

	data, err := h.manager.Login(c.UserContext(), req)
	if err != nil {
		logger.Error(c.UserContext(), c.Path(), err.Error())
		if errors.Is(err, constant.ErrInvalidIdentifierOrPassword) {
			return response.Error(c, constant.ErrCodeInvalidParam, err.Error())
		}
		return response.Error(c, constant.ErrCodeInternalServer, err.Error())
	}

	return response.Success(c, req, data)
}
