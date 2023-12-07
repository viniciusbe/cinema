package entities

import (
	"time"

	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Description string `json:"description"`
	Films       []Film `gorm:"many2many:film_genders;"`
}

type Director struct {
	gorm.Model
	Name  string `json:"name"`
	Films []Film
}

type Film struct {
	gorm.Model
	Name       string `json:"name"`
	Poster     string `json:"poster"`
	Duration   uint   `json:"duration"`
	Synopsis   string `json:"synopsis"`
	Age        uint   `json:"age"`
	DirectorID uint
	Director   Director  `json:"director"`
	Genders    []Gender  `gorm:"many2many:film_genders;" json:"genders"`
	Sessions   []Session `json:"sessions"`
}

type Session struct {
	gorm.Model
	Time     time.Time `json:"time"`
	Language string    `json:"language"`
	Room     uint      `json:"room"`
	FilmID   uint      `json:"filmId"`
	Film     Film
	Tickets  []Ticket
}

type Buyer struct {
	gorm.Model
	Name     string   `json:"name"`
	Document string   `json:"document"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Tickets  []Ticket `json:"tickets"`
}

type Ticket struct {
	gorm.Model
	Seat      string `json:"seat"`
	Modality  string `json:"modality"`
	BuyerID   uint   `json:"buyerId"`
	SessionID uint   `json:"sessionId"`
	Buyer     Buyer
	Session   Session `json:"session"`
}
