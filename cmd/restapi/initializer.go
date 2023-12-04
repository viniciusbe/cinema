package main

import (
	"cinema/internal/core/services/buyerserv"
	"cinema/internal/core/services/filmserv"
	"cinema/internal/core/services/sessionserv"
	"cinema/internal/core/services/ticketserv"
	"cinema/internal/repositories/buyerrepo"
	"cinema/internal/repositories/filmrepo"
	"cinema/internal/repositories/sessionrepo"
	"cinema/internal/repositories/ticketrepo"

	"gorm.io/gorm"
)

func Buyer(db *gorm.DB) *buyerserv.Service {
	repo := buyerrepo.NewGormRepository(db)
	return buyerserv.New(repo)
}

func Film(db *gorm.DB) *filmserv.Service {
	repo := filmrepo.NewGormRepository(db)
	return filmserv.New(repo)
}

func Session(db *gorm.DB) *sessionserv.Service {
	repo := sessionrepo.NewGormRepository(db)
	return sessionserv.New(repo)
}

func Ticket(db *gorm.DB) *ticketserv.Service {
	repo := ticketrepo.NewGormRepository(db)
	return ticketserv.New(repo)
}
