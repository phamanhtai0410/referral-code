package main

import (
	"os"
	"os/signal"
	"syscall"

	"log"

	"example.com/refcode/v1/app"
	"example.com/refcode/v1/app/tasks"
	"example.com/refcode/v1/pkg/configs"
	"example.com/refcode/v1/pkg/workers"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/urfave/cli"

	_ "example.com/refcode/v1/docs" // load API Docs files (Swagger)
	"example.com/refcode/v1/pkg/middleware"
	"example.com/refcode/v1/pkg/routes"
	// _ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

var (
	client *cli.App
)

func init() {
	// Initialise a CLI app
	client = cli.NewApp()
	client.Name = "machinery"
	client.Usage = "machinery worker and handler reference code"
	client.Version = "0.0.0"
}

func startServer(cfg *configs.Worker) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// create a new app
	_app := app.New()
	_app.Shutdown(sigChan)

	// register background tasks
	_app.BackgroundTask(
		tasks.RefCodeCounterUsed,
		tasks.ListenChainEvent,
	)

	_app.Worker(cfg)

	// register middleware
	_app.Middleware(
		middleware.FiberMiddleware,
	)

	// register route
	_app.Route(
		routes.HealthCheck,
		routes.RefCodeRoute,
		routes.SwaggerRoute,
		routes.NotFoundRoute,
	)
	_app.Run()
}

func main() {

	cnf := &configs.Worker{
		Config: &config.Config{
			Broker:          configs.BrokerUrl,
			DefaultQueue:    "machinery_tasks",
			ResultBackend:   configs.CacheUrl,
			ResultsExpireIn: 3600,
			AMQP: &config.AMQPConfig{
				Exchange:      "machinery_exchange",
				ExchangeType:  "direct",
				BindingKey:    "machinery_task",
				PrefetchCount: 3,
			},
		},
		Task: map[string]interface{}{
			"Worker.HealthCheck":      tasks.HealthCheck,
			"Worker.SaveCodeUsed":     tasks.SaveRefCodeUsed,
			"Worker.SaveHistory":      tasks.SaveHistory,
			"Worker.UpdateUserRecord": tasks.UpdateUserRecord,
			"Worker.SaveUserInfo":     tasks.SaveUserInfo,
		},
	}

	client.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Command.Name)
				if err := workers.Execute(cnf, "consume", 12); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "send example tasks ",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Command.Name)
				startServer(cnf)
				return nil
			},
		},
	}

	// Run the CLI app
	client.Run(os.Args)
}
