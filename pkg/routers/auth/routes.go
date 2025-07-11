package auth

import (
	"github.com/gofiber/fiber/v2"
	"mongo_db/app/controllers"
)

func Route(app *fiber.App) {
	route := app.Group("/auth")
	route.Post("/login", controllers.Login)
	route.Delete("/delete/:id", controllers.DeleteUser)
	route.Post("/email", controllers.EmailVerify)
	route.Get("/email/verify", controllers.TokenVerification)
	route.Get("/users", controllers.GetUsers)
}
