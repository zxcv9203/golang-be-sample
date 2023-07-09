package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zxcv9203/golang-be-sample/internal/service"
	"github.com/zxcv9203/golang-be-sample/internal/transport/rest/request"
	"strconv"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := new(request.Post)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	response := h.service.Save(req)

	return c.JSON(response)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	req := new(request.Post)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.Update(id, req)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *Handler) FindById(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.FindById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(response)
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = h.service.DeleteById(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return nil
}

func (h *Handler) Find(c *fiber.Ctx) error {
	pageRequest, err := parsePage(c.Query("size"), c.Query("page"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	response := h.service.Find(pageRequest)

	return c.JSON(response)
}

func parsePage(size string, page string) (request.Page, error) {
	parseSize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		parseSize = 20
	}
	parsePage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		parsePage = 1
	}

	pageRequest := request.NewPageRequest(parseSize, parsePage)
	return pageRequest, nil
}
