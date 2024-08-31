package response

import (
	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/proto"
)

type response struct {
	Data     interface{} `json:"data,omitempty"`
	Error    *int32      `json:"error,omitempty"`
	ErrorMsg *string     `json:"error_msg,omitempty"`
}

func Success(c *fiber.Ctx, req, data interface{}) error {
	res := response{
		Data: data,
	}
	logger.Data(c.UserContext(), c.Path(), req, data)
	return c.Status(fiber.StatusOK).JSON(res)
}

func Error(c *fiber.Ctx, req interface{}, errCode int32, errMsg string) error {
	res := response{
		Error:    proto.Int32(errCode),
		ErrorMsg: proto.String(errMsg),
	}
	logger.Data(c.UserContext(), c.Path(), req, res)
	return c.Status(fiber.StatusOK).JSON(res)
}
