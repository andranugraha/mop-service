package middleware

import (
	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/common/tracing"
	"github.com/gofiber/fiber/v2"
)

func ServiceLoggingMiddleware(c *fiber.Ctx) error {
	ctx := c.UserContext()

	defer func() {
		if r := recover(); r != nil {
			logger.Panic(ctx, "panic: %v", r)
		}
	}()

	ctx = tracing.NewContextWithTracingID(ctx)
	c.SetUserContext(ctx)

	logger.Info(ctx, c.Path(), "request: %v", string(c.Request().Body()))

	err := c.Next()
	if err != nil {
		return err
	}

	return nil
}
