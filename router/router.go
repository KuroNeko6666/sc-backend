package router

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	base := app.Group("/api")
	admin(base)
	user(base)
}
