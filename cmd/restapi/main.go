package main

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/handlers/gormdb"
	"cinema/internal/handlers/rest/buyerrest"
	"cinema/internal/handlers/rest/film"
	"cinema/internal/handlers/rest/session"
	"cinema/internal/handlers/rest/ticket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	db, err := gormdb.SetupGormDB()

	if err != nil {

		panic(err)
	}

	db.AutoMigrate(&entities.Gender{}, &entities.Director{}, &entities.Film{}, &entities.Session{}, &entities.Buyer{}, &entities.Ticket{})

	buyerService := Buyer(db)
	buyerRoute := app.Group("/buyers")
	buyerrest.Route(buyerRoute, buyerService)

	filmService := Film(db)
	filmRoute := app.Group("/films")
	film.Route(filmRoute, filmService)

	sessionService := Session(db)
	sessionRoute := app.Group("/sessions")
	session.Route(sessionRoute, sessionService)

	ticketService := Ticket(db)
	ticketRoute := app.Group("/tickets")
	ticket.Route(ticketRoute, ticketService)

	app.Listen(":3000")
}
