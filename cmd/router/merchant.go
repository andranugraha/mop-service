package router

import (
	"github.com/empnefsi/mop-service/internal/handler/merchant"
	"github.com/gofiber/fiber/v2"
)

func registerMerchantRoutes(router fiber.Router) {
	merchantHandler := merchant.NewHandler()
	merchantRouter := router.Group("/merchant")
	merchantRouter.Get("/:merchant_id/payment-types", merchantHandler.GetMerchantActivePaymentTypes)
	merchantRouter.Get("/:merchant_id/additional-fees", merchantHandler.GetMerchantActiveAdditionalFees)
}
