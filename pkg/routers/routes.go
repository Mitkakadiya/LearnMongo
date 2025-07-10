package routers

import (
	"github.com/gofiber/fiber/v2"
	"mongo_db/pkg/routers/auth"
)

func Router(app *fiber.App) {
	auth.Route(app)
}
