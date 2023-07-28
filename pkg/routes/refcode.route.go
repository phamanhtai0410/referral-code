package routes

import (
	"example.com/refcode/v1/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RefCodeRoute(a *fiber.App) {
	// route := a.Group("/refcode")
	a.Get("/get", controllers.RefCodeGenerate)
	a.Post("/save/", controllers.SaveRefCodeUsed)
	a.Get("/tracking/:address", controllers.RefCodeTracking)
	//a.Get("/test", controllers.Test)
}
