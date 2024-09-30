package router

import (
	"github.com/empnefsi/mop-service/cmd/middleware"
	"github.com/empnefsi/mop-service/internal/handler/order"
	"github.com/gofiber/fiber/v2"
)

func registerOrderRoutes(router fiber.Router) {
	orderHandler := order.NewHandler()
	orderRouter := router.Group("/order")
	orderRouter.Post("", orderHandler.CreateOrder)
	orderRouter.Post("/payment-callback", orderHandler.PaymentCallback)
	orderRouter.Get("/:order_id/push-payment-event", orderHandler.PushPaymentEvent)

	orderRouter.Post("/pay", middleware.AuthMiddleware, middleware.CashierMiddleware, orderHandler.PayOrder)
}
