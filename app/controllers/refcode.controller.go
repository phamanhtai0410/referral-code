package controllers

import (
	"log"

	"example.com/refcode/v1/app/services"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

// RefCodeGenerate Generate RefCode.
// @Description Generate RefCode.
// @Tags RefCode
// @Accept json
// @Produce json
// @Success 200 {object} models.RefCode
// @Router /refcode/gen [GET]
func RefCodeGenerate(c *fiber.Ctx) error {
	address := c.Query("address", "NONE")
	if address == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    nil,
			"msg":     "address is required",
			"success": false,
		})
	}
	code := services.GenRefCode(address)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"msg":     "referral code generated",
		"success": true,
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
// @Success 200 {object} models.RefCode
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
	services.SaveRefCodeInfo(request)
	return c.JSON(fiber.Map{
		"msg":     "ok",
		"success": true,
	})
}

// func RefCodeCounter(c *fiber.Ctx) error {
// 	tableCodeUsed := new(models.RefCodeUsed)
// 	tableCodeUsed.RefCode = utils.String2Int64(c.Query("code"))
// 	currentTime := time.Now()
// 	oneMinuteAgo := currentTime.Add(-2 * time.Minute)
// 	count, _ := tableCodeUsed.CountDocumentsByTime(
// 		oneMinuteAgo.UTC(),
// 		currentTime.UTC(),
// 	)
// 	return c.JSON(fiber.Map{
// 		"msg":     count,
// 		"success": true,
// 	})
// }
