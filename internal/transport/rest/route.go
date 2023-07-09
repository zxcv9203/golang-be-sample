package rest

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, handler *Handler) {
	app.Post("/posts", handler.Create)
	app.Put("/posts/:id", handler.Update)
	app.Get("/posts/:id", handler.FindById)
}
