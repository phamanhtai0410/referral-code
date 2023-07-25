package app

import (
	"fmt"
	"log"
	"os"

	"example.com/refcode/v1/pkg/configs"
	"example.com/refcode/v1/pkg/constants"
	"example.com/refcode/v1/pkg/utils"
	"example.com/refcode/v1/platform/cache"
	"example.com/refcode/v1/platform/database"
	"example.com/refcode/v1/platform/queue"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

type _App struct {
	engine *fiber.App
}

type IApp interface {
	Tasks(...worker)
	Middleware(...middleware) IApp
	Route(...route) IApp
	Run()
	Shutdown(<-chan os.Signal)
}

type middleware func(*fiber.App)
type route func(*fiber.App)
type worker func(*amqp.Delivery)

func New() IApp {
	config := configs.FiberConfig()
	return &_App{
		engine: fiber.New(config),
	}
}

func (app *_App) Tasks(workers ...worker) {
	msgs, err := queue.Consume(constants.WorkerQueue)
	if err != nil {
		log.Fatal("failed to consume refcode info: ", err)
	}
	go func() {
		for msg := range msgs {
			for _, worker := range workers {
				go worker(&msg)
			}
		}
	}()
}

func (app *_App) Middleware(middlewares ...middleware) IApp {
	for _, middleware := range middlewares {
		middleware(app.engine)
	}
	return app
}

func (app *_App) Route(routes ...route) IApp {
	for _, route := range routes {
		route(app.engine)
	}
	return app
}

func (app *_App) Shutdown(sig <-chan os.Signal) {
	go func() {
		<-sig
		fmt.Println()
		queue.Shutdown()
		database.Shutdown()
		cache.Shutdown()
		if configs.StageStatus == "prod" {
			log.Println("[SERVER] Server is shutting down ..")
			if err := app.engine.Shutdown(); err != nil {
				log.Printf("Oops... Server is not shutting down! Reason: %v", err)
			}
		} else {
			os.Exit(0)
		}
	}()
}

func (app *_App) Run() {
	utils.StartServer(app.engine)
}
