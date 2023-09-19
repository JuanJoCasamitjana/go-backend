package setup

import (
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func SetUpAndRun() {
	//Initialize app
	app := fiber.New()
	routes.SetUpRoutes(app)
	//Config

	//Start
	app.Listen(":8080")
}
