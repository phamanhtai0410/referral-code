package controllers

import (
	"example.com/refcode/v1/app/tasks"
	"example.com/refcode/v1/pkg/workers"
	"github.com/gofiber/fiber/v2"
)

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

func HealthCheckWorker(c *fiber.Ctx) error {
	err := workers.Delay("Worker.HealthCheck", tasks.HealthCheck, int64(1))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "failed",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
