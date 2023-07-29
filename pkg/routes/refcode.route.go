package routes

import (
	"example.com/refcode/v1/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RefCodeRoute(a *fiber.App) {
	//a.Get("/test", controllers.Test)
	a.Post("/save/", controllers.SaveRefCodeUsed)
	a.Get("/tracking/:domain", controllers.RefCodeTracking)
	a.Get("/history/:domain", controllers.WithdrawHistory)

}
