package controllers

import "github.com/gofiber/fiber/v2"

// HealthCheck health check api server.
// @Description health check api server.
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 
// @Router /healthcheck [GET]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
