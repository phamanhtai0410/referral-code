package routes

import (
	"example.com/refcode/v1/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(a *fiber.App) {
	a.Get("/healthcheck", controllers.HealthCheck)
	a.Get("/workercheck", controllers.HealthCheckWorker)
}
