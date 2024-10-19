package landing

import (
	"github.com/empnefsi/mop-service/internal/common/constant"
	"github.com/empnefsi/mop-service/internal/common/response"
	"github.com/empnefsi/mop-service/internal/dto/landing"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get Landing Banners
// @Description Get Landing Banners
// @Tags Landing
// @Accept json
// @Produce json
// @Param code path string true "Code"
// @Success 200 {object} landing.GetLandingBannersResponse
// @Router /api/v1/landing/{code}/banners [get]
func (h *impl) GetLandingBanners(c *fiber.Ctx) error {
	code := c.Params("code")
	req := &landing.LandingRequest{Code: code}
	if code == "" {
		return response.Error(c, req, constant.ErrCodeInvalidParam, "Code is required")
	}

	data, err := h.manager.GetLandingBanners(c.UserContext(), code)
	if err != nil {
		return response.Error(c, req, constant.GetErrorCode(err), err.Error())
	}

	return response.Success(c, req, data)
}
