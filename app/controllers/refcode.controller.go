package controllers

import (
	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/app/services"
	"example.com/refcode/v1/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"log"
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
		"msg": "ok",
		"data": fiber.Map{
			"count": resp.Count,
			"rate":  resp.Rate,
			"level": resp.Level,
		},
		"success": true,
	})
}

//func Test(c *fiber.Ctx) error {
//	address := common.HexToAddress(configs.ContractAddr)
//	refCode, _ := contract.NewRefCode(address, blockchain.GetClient())
//	//addr := common.HexToAddress("0x8f9d9aA7B313cf9360d4E61D1Ae809443f97aCad")
//	//refCode.UploadSingleData(blockchain.NewTransactOpts(&configs.TransactOptsConfig{
//	//	ChainID:    11155111,
//	//	PrivateKey: "5d935ba1a8c41db1f4c5aa8db93a9e0d429d9832b1004623a93e197df65d7ab2",
//	//}), addr, big.NewInt(0.1*1e18))
//	//refCode.Withdraw(blockchain.NewTransactOpts(&configs.TransactOptsConfig{
//	//	ChainID:    11155111,
//	//	PrivateKey: "5d935ba1a8c41db1f4c5aa8db93a9e0d429d9832b1004623a93e197df65d7ab2",
//	//}), big.NewInt(0.05*1e18))
//	//return c.JSON(fiber.Map{
//	//	"success": true,
//	//})
//}
