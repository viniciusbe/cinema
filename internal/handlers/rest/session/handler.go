package session

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/sessionserv"
	"cinema/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *sessionserv.Service
}

func (h *Handler) ListAll(c *fiber.Ctx) error {

	sessions, err := h.service.ListAll()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(sessions)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	session := new(entities.Session)
	c.BodyParser(session)
	err := h.service.Create(session)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(session)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	session := new(entities.Session)
	c.BodyParser(session)

	session.ID = utils.StringToUint(id)

	err := h.service.Update(session)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(session)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(400).SendString("Não existe")
	}

	return c.SendString("Excluido com sucesso")
}
