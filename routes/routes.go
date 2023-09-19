package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Hello)
}
