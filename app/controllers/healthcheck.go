package controllers

import (
	"example.com/refcode/v1/app/tasks"
	"example.com/refcode/v1/pkg/configs"
	"example.com/refcode/v1/pkg/workers"
	"github.com/RichardKnop/machinery/v1/config"
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
	cnf := &configs.Worker{
		Config: &config.Config{
			Broker:          "amqp://guest:guest@localhost:5672/",
			DefaultQueue:    "machinery_tasks",
			ResultBackend:   "redis://localhost:6379/0",
			ResultsExpireIn: 3600,
			AMQP: &config.AMQPConfig{
				Exchange:      "machinery_exchange",
				ExchangeType:  "direct",
				BindingKey:    "machinery_task",
				PrefetchCount: 3,
			},
		},
		Task: map[string]interface{}{
			"healthcheck": tasks.HealthCheck,
		},
	}
	workers.StartServer(cnf)
	workers.Delay("healthcheck", tasks.HealthCheck, int64(1))
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
