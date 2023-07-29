package controllers

import (
	"log"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/app/services"
	"example.com/refcode/v1/pkg/constants"
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
	request := new(schemas.RefCodeUsedRequest)
	if err := c.BodyParser(request); err != nil {
		log.Printf("Error parsing request body: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     constants.ErrorBodyParser,
			"success": false,
		})
	}
	if err := services.SaveRefCodeInfo(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     err.Error(),
			"success": false,
		})
	}
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
	domain := c.Params("domain", "NONE")
	if domain == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"msg":     "domain is required",
			"success": false,
		})
	}
	resp, err := services.RefCodeTracking(domain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     err.Error(),
			"data":    nil,
			"success": false,
		})
	}
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"data":    resp,
		"success": true,
	})
}

func WithdrawHistory(c *fiber.Ctx) error {
	domain := c.Params("domain", "NONE")
	if domain == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"msg":     "domain is required",
			"success": false,
		})
	}
	resp, err := services.WithdrawHistory(domain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     err.Error(),
			"data":    nil,
			"success": false,
		})
	}
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"data":    resp,
		"success": true,
	})
}
