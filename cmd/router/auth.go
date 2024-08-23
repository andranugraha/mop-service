package router

import (
	"github.com/empnefsi/mop-service/internal/handler/auth"
	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(router fiber.Router) {
	authHandler := auth.NewHandler()
	authRouter := router.Group("/auth")
	authRouter.Post("/login", authHandler.Login)
}
