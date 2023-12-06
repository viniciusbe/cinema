package ticket

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/ticketserv"
	"cinema/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *ticketserv.Service
}

func (h *Handler) ListAll(c *fiber.Ctx) error {

	tickets, err := h.service.ListAll()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(tickets)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	ticket := new(entities.Ticket)
	c.BodyParser(ticket)
	err := h.service.Create(ticket)

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(ticket)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	ticket := new(entities.Ticket)
	c.BodyParser(ticket)

	ticket.ID = utils.StringToUint(id)

	err := h.service.Update(ticket)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(ticket)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(400).SendString("Não existe")
	}

	return c.SendString("Excluido com sucesso")
}
