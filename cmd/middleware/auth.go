package middleware

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/context"
	"github.com/empnefsi/mop-service/internal/common/strings"
	"github.com/empnefsi/mop-service/internal/module/user"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check if user is authenticated
	token := c.Get("Authorization")
	claims, err := strings.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":     constant.ErrCodeUnauthorized,
			"error_msg": constant.ErrUnauthorized.Error(),
		})
	}

	// Set user context
	ctx := c.UserContext()
	ctx = context.WithValue(ctx, "user", claims)

	c.SetUserContext(ctx)

	// Continue stack
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	// Check if user is an admin
	role := c.UserContext().Value("role").(uint32)
	if role != user.RoleAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":     constant.ErrCodeUnauthorized,
			"error_msg": constant.ErrUnauthorized.Error(),
		})
	}

	// Continue stack
	return c.Next()
}

func CashierMiddleware(c *fiber.Ctx) error {
	// Check if user is a cashier
	role := contextutil.GetRole(c.UserContext())
	if role != user.RoleCashier {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":     constant.ErrCodeUnauthorized,
			"error_msg": constant.ErrUnauthorized.Error(),
		})
	}

	// Continue stack
	return c.Next()
}
