package order

import (
	orderManager "github.com/empnefsi/mop-service/internal/manager/order"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateOrder(c *fiber.Ctx) error
	PayOrder(c *fiber.Ctx) error
}

type impl struct {
	manager orderManager.Manager
}

func NewHandler() Handler {
	return &impl{
		manager: orderManager.NewManager(),
	}
}
