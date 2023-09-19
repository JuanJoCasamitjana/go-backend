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
	routes.SetUpRoutes(app)
	//Config

	//Start
	app.Listen(":8080")
}
