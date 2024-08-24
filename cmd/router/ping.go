package router

import (
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/gofiber/fiber/v2"
)

func registerPingRoutes(fiberRouter fiber.Router) {
	fiberRouter.Get("/ping", func(c *fiber.Ctx) error {
		return response.Success(c, nil, "pong")
	})
}
