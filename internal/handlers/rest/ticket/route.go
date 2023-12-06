package ticket

import (
	"cinema/internal/core/services/ticketserv"

	"github.com/gofiber/fiber/v2"
)

func Route(route fiber.Router, service *ticketserv.Service) {
	handler := Handler{
		service: service,
	}

	route.Get("/", handler.ListAll)
	route.Post("/", handler.Create)
	route.Put("/:id", handler.Edit)
	route.Delete("/:id", handler.Delete)
}
