package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxcv9203/golang-be-sample/internal/service"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) CreateHandler(c *fiber.Ctx) error {
	req := new(request.Post)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	response := h.service.Save(req)

	return c.JSON(response)
}
