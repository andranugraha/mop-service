package landing

import (
	"github.com/empnefsi/mop-service/internal/manager/landing"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Landing(c *fiber.Ctx) error
}

type impl struct {
	manager landing.Manager
}

func NewHandler() Handler {
	return &impl{
		manager: landing.NewManager(),
	}
}
