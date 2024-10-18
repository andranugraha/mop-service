package router

import (
	_ "github.com/empnefsi/mop-service/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/metrics", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := app.Group("/api/v1")
	registerPingRoutes(v1)
	registerAuthRoutes(v1)
	registerLandingRoutes(v1)
	registerOrderRoutes(v1)
	registerMerchantRoutes(v1)
}
