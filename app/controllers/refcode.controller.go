package controllers

import (
	"log"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/app/services"
	"example.com/refcode/v1/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

// RefCodeGenerate Generate RefCode.
// @Description Generate RefCode.
// @Tags RefCode
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /refcode/gen [GET]
func RefCodeGenerate(c *fiber.Ctx) error {
	address := c.Query("address", "NONE")
	domain := c.Query("domain", "NONE")
	if address == "NONE" || domain == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    nil,
			"msg":     "address, domain is required",
			"success": false,
		})
	}
	id, err := services.ReferralCodeHandle(&schemas.RefCodeRequest{
		Address: address,
		Domain:  domain,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    nil,
			"msg":     err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.RefCodeResponse{
		Id:      id,
		Code:    domain,
		Message: "referral code generated",
		Success: true,
	})
}

// SaveRefCodeUsed Save to Database.
// @Description Save RefCode to Database.
// @Tags User
// @Accept json
// @Produce json
// @Param refcode body number true "RefCode"
// @Param domain body string true "Domain"
// @Param price body number true "Price"
// @Success 200 {object} models.CodeUsed
// @Router /refcode/save [POST]
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
	address := c.Params("address", "NONE")
	if address == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"msg":     "address is required",
			"success": false,
		})
	}
	resp, err := services.RefCodeTracking(address)
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

//func Test(c *fiber.Ctx) error {
//	codeUsedModel := new(models.CodeUsed)
//	records, err := codeUsedModel.GetDocumentsByTime(
//		oneMinuteAgo.UTC(),
//		currentTime.UTC(),
//	)
//}
