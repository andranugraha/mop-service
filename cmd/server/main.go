package main

import (
	"github.com/empnefsi/mop-service/cmd/middleware"
	"github.com/empnefsi/mop-service/cmd/router"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	run()
}

func run() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cache.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Use(middleware.ServiceLoggingMiddleware)

	router.RegisterRoutes(app)

	port := config.GetPort()
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
