package buyerrest

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/buyerserv"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *buyerserv.Service
}

func (h *Handler) ListAll(c *fiber.Ctx) error {

	buyers, err := h.service.ListAll()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(buyers)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	buyer := new(entities.Buyer)
	c.BodyParser(buyer)
	err := h.service.Create(buyer)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(buyer)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	buyer := new(entities.Buyer)
	c.BodyParser(buyer)

	u64ID, _ := strconv.ParseUint(id, 10, 32)
	buyer.ID = uint(u64ID)

	err := h.service.Update(buyer)

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(buyer)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(400).SendString("Não existe")
	}

	return c.SendString("Excluido com sucesso")
}

func (h *Handler) GetDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	buyer, tickets, err := h.service.Get(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	buyer.Tickets = tickets

	return c.JSON(buyer)
}
