package router

import (
	"github.com/empnefsi/mop-service/internal/handler/landing"
	"github.com/gofiber/fiber/v2"
)

func registerLandingRoutes(router fiber.Router) {
	landingHandler := landing.NewHandler()
	landingRouter := router.Group("/landing")
	landingRouter.Get("/:code", landingHandler.Landing)
}
