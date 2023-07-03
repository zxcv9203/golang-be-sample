package rest

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, handler *Handler) {

	app.Post("/posts", handler.CreateHandler)
	app.Put("/posts/:id", handler.UpdateHandler)
}
