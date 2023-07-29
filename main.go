package main

import (
	"os"
	"os/signal"
	"syscall"

	"example.com/refcode/v1/app/tasks"
	"example.com/refcode/v1/app/workers"

	"example.com/refcode/v1/app"
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
func main() {

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

	// register workers
	_app.Worker(
		workers.SaveRefCodeUsed,
	)

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

	// run application
	_app.Run()

}
