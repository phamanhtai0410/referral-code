package routes

import (
	"example.com/refcode/v1/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RefCodeRoute(a *fiber.App) {
	route := a.Group("/refcode")
	route.Get("/get", controllers.RefCodeGenerate)
	route.Post("/save", controllers.SaveRefCodeUsed)
	//route.Get("/share", middleware.Auth, controllers.RefCodeCounter)
}
