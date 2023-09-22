package entities

import (
	"time"

	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Description string
	Films       []Film `gorm:"many2many:film_genders;"`
}

type Director struct {
	gorm.Model
	Name  string
	Films []Film
}

type Film struct {
	gorm.Model
	Name       string
	Duration   uint
	Synopsis   string
	Age        uint
	DirectorID uint
	Director   Director
	Genders    []Gender `gorm:"many2many:film_genders;"`
	Sessions   []Session
}

type Session struct {
	gorm.Model
	Time     time.Time
	Language string
	Room     uint
	FilmID   uint
	Film     Film
	Tickets  []Ticket
}

type Buyer struct {
	gorm.Model
	Name     string
	Document string
	Tickets  []Ticket
}

type Ticket struct {
	gorm.Model
	Seat      string
	Modality  string
	BuyerID   uint
	SessionID uint
	Buyer     Buyer
	Session   Session
}
