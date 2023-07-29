package routes

import (
	"example.com/refcode/v1/app/controllers"
	"example.com/refcode/v1/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func RefCodeRoute(a *fiber.App) {
	//a.Get("/test", controllers.Test)
	a.Post("/save", middleware.ValidateData, controllers.SaveRefCodeUsed)
	a.Get("/tracking/:domain", middleware.ValidateDomain, controllers.RefCodeTracking)
	a.Get("/history/:domain", middleware.ValidateDomain, controllers.WithdrawHistory)

}
