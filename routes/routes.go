package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Hello)
	app.Get("/todos", handlers.GetAllTodos)
	app.Get("/finished", handlers.GetCompletedAllTodos)
	app.Get("/unfinished", handlers.GetUncompletedAllTodos)
	app.Post("/todos", handlers.CreateToDo)
	app.Put("/todos", handlers.UpdateTodo)
	app.Put("/todos/:id", handlers.UpdateTodoById)
	app.Delete("/todos", handlers.DeleteTodo)
	app.Delete("/todos/:id", handlers.DeleteTodoById)
}
