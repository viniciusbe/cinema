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
	DirectorID string
	Films      []Film
}

type Film struct {
	Name     string
	Duration string
	Synopsis string
	Age      string
	FilmID   string
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
	BuyerID  string
	Tickets  []Ticket
}

type Ticket struct {
	Seat     string
	Modality string
	TickedID uint
	Buyer    Buyer
	Session  Session
}
