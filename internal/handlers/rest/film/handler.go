package film

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/services/filmserv"
	"cinema/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *filmserv.Service
}

func (h *Handler) ListAll(c *fiber.Ctx) error {

	films, err := h.service.ListAll()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(films)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	film := new(entities.Film)
	c.BodyParser(film)
	err := h.service.Create(film)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(film)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	film := new(entities.Film)
	c.BodyParser(film)

	film.ID = utils.StringToUint(id)

	err := h.service.Update(film)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(film)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(400).SendString("Não existe")
	}

	return c.SendString("Excluido com sucesso")
}
