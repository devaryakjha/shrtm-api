package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()
	app.Get("/", greet)
	adaptor.FiberApp(app)(w, r)
}

func greet(c *fiber.Ctx) error {
	c.Response().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.SendString("Hello World 👋!")
}
