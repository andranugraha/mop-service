package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func RegisterRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/metrics", monitor.New())

	v1 := app.Group("/api/v1")
	registerPingRoutes(v1)
	registerAuthRoutes(v1)
}
