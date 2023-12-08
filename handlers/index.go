package handlers

import "github.com/gofiber/fiber/v2"

func SendIndex(c *fiber.Ctx) error {
	return c.SendFile("static/front-end/dist/index.html")
}
