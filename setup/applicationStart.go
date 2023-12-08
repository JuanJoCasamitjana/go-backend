package setup

import (
	"backend/database"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetUpAndRun() {
	//Database
	database.InitDB()
	err := database.CreateTables()
	if err != nil {
		log.Fatal(err)
	}
	//Initialize app
	app := fiber.New()
	app.Static("/assets", "static/front-end/dist/assets")
	app.Static("/favicon.ico", "static/front-end/dist/favicon.ico")
	routes.SetUpRoutes(app)
	//Config

	//Start
	app.Listen(":8080")
}
