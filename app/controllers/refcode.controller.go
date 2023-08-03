package controllers

import (
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/app/services"
	"example.com/refcode/v1/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// SaveRefCodeUsed Save to Database.
// @Description Save RefCode to Database.
// @Tags User
// @Accept json
// @Produce json
// @Param refcode body number true "RefCode"
// @Param domain body string true "Domain"
// @Param price body number true "Price"
// @Success 200 {object} models.CodeUsed
// @Router /save [POST]
func SaveRefCodeUsed(c *fiber.Ctx) error {
	req, ok := c.Locals("/save").(schemas.CodeUsedRequest)
	if !ok {
		// Handle error (e.g., return an error response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "Failed to get parsed data from context",
			"success": false,
		})
	}
	price, err := utils.String2Float64(req.Price)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     err.Error(),
			"success": false,
		})
	}
	go services.SaveRefCodeInfo(&schemas.RefCodeUsedRequest{
		ReferralCode: req.ReferralCode,
		Price:        price,
		Address:      req.Address,
		Domain:       req.Address,
	})
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"success": true,
	})
}

// RefCodeTracking Tracking code
// @Description Tracking code.
// @Tags Tracking
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /tracking/:address [GET]
func RefCodeTracking(c *fiber.Ctx) error {
	domain, ok := c.Locals("domain").(string)
	if !ok {
		// Handle error (e.g., return an error response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"msg":     "Failed to get parsed data from context",
			"success": false,
		})
	}
	resp, err := services.RefCodeTracking(domain)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg": "ok",
			"data": schemas.TrackingResponse{
				ReferralCode: domain,
				Count:        0,
				Level:        "",
				Rate:         0,
				TotalEarn:    0,
			},
			"success": true,
		})
	}
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"data":    resp,
		"success": true,
	})
}

func WithdrawHistory(c *fiber.Ctx) error {
	domain, ok := c.Locals("domain").(string)
	if !ok {
		// Handle error (e.g., return an error response)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "Failed to get parsed data from context",
			"success": false,
		})
	}
	resp, err := services.WithdrawHistory(domain)
	if err != nil {
		var status int
		data := make([]interface{}, 0)
		if err.Error() == "no data" {
			status = fiber.StatusOK
		} else {
			status = fiber.StatusBadRequest
			data = nil
		}
		return c.Status(status).JSON(fiber.Map{
			"msg":     err.Error(),
			"data":    data,
			"success": false,
		})
	}
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"data":    resp,
		"success": true,
	})
}
