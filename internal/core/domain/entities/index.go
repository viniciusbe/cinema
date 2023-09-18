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
	Duration   uint16
	Synopsis   string
	Age        uint8
	DirectorID uint
	Genders    []Gender `gorm:"many2many:film_genders;"`
	Sessions   []Session
}

type Session struct {
	gorm.Model
	Time     time.Time
	Language string
	Room     uint8
	FilmID   uint
	Tickets  []Ticket
}

type Ticket struct {
	gorm.Model
	Seat      string
	Modality  string
	BuyerID   uint
	SessionID uint
}

type Buyer struct {
	gorm.Model
	Name     string
	Document string
	Tickets  []Ticket
}