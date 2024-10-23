package handlers

import (
	"github.com/Ashmn07/ticket-booking/models"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	return nil
}
func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}
func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}
func (h *TicketHandler) UpdateOne(ctx *fiber.Ctx) error {
	return nil
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:eventId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:eventId", handler.UpdateOne)
}
