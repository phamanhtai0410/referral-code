package middleware

import (
	"log"

	"example.com/refcode/v1/app/schemas"
	"example.com/refcode/v1/pkg/constants"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func ValidateDomain(c *fiber.Ctx) error {
	domain := c.Params("domain", "NONE")
	if domain == "NONE" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    nil,
			"msg":     "domain is required",
			"success": false,
		})
	}
	c.Locals("domain", domain)
	return c.Next()
}

func ValidateData(c *fiber.Ctx) error {
	validate := validator.New()
	switch c.Path() {
	case "/save":
		request := new(schemas.RefCodeUsedRequest)
		if err := c.BodyParser(request); err != nil {
			log.Printf("Error parsing request body: " + err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":     constants.ErrorBodyParser,
				"success": false,
			})
		}
		if err := validate.Struct(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":     "missing required fields",
				"success": false,
			})
		}
		c.Locals("/save", *request)
	}
	return c.Next()
}
