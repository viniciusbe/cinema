package entities

import (
	"time"
)

type Gender struct {
	Description string
	GenderID    string
	Films       []Film
}

type Director struct {
	Name       string
	DirectorID uint
	Films      []Film
}

type Film struct {
	Name     string
	Duration uint
	Synopsis string
	Age      uint
	FilmID   uint
	Director Director
	Genders  []Gender
	Sessions []Session
}

type Session struct {
	Time      time.Time
	Language  string
	Room      uint
	SessionID uint
	Film      Film
	Tickets   []Ticket
}

type Buyer struct {
	Name     string
	Document string
	BuyerID  uint
	Tickets  []Ticket
}

type Ticket struct {
	Seat     string
	Modality string
	TickedID uint
	Buyer    Buyer
	Session  Session
}
