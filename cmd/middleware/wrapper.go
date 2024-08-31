package middleware

import (
	"context"
	"encoding/json"
	"time"

	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/common/tracing"
	"github.com/empnefsi/mop-service/internal/config"
	"github.com/gofiber/fiber/v2"
)

func TrafficWrapperMiddleware(c *fiber.Ctx) error {
	ctx := c.UserContext()
	defer func() {
		if r := recover(); r != nil {
			logger.Panic(ctx, "panic: %v", r)
		}
	}()

	ctx = tracing.NewContextWithTracingID(ctx)
	tracingId := tracing.GetTracingIDFromCtx(ctx)
	c.Set("X-Request-ID", tracingId)

	startTime := time.Now()
	ctx = tracing.AppendMetadataToIncomingContext(ctx, "start_time", startTime.Format(time.RFC3339Nano))

	c.SetUserContext(ctx)

	timeout := config.GetTimeout()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if timeoutCtx.Err() == nil {
					logger.Panic(ctx, "panic: %v", r)
				}
				done <- nil
			}
			done <- nil
		}()
		done <- c.Next()
	}()

	select {
	case <-timeoutCtx.Done():
		errMessage := "request timed out"
		if ctxErr := timeoutCtx.Err(); ctxErr != nil {
			errMessage = ctxErr.Error()
		}

		var reqBody interface{}
		_ = json.Unmarshal(c.Request().Body(), &reqBody)

		return response.Error(c, reqBody, constant.ErrCodeTimeout, errMessage)
	case err := <-done:
		if err != nil {
			return err
		}
	}

	return nil
}
