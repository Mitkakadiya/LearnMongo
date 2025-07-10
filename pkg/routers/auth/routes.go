package auth

import (
	"github.com/gofiber/fiber/v2"
	"mongo_db/app/controllers"
)

func Route(app *fiber.App) {
	route := app.Group("auth")
	route.Post("/login", controllers.Login)
}
