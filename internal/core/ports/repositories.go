package ports

import (
	"cinema/internal/core/domain/entities"
	"time"
)

type GenderRepository interface {
	ListAll() ([]entities.Gender, error)
	Find(id string) (*entities.Gender, error)
	Insert(gender *entities.Gender) error
	Save(gender *entities.Gender) error
	Delete(id string) error
}

type DirectorRepository interface {
	ListAll() ([]entities.Director, error)
	Find(id string) (*entities.Director, error)
	Insert(director *entities.Director) error
	Save(director *entities.Director) error
	Delete(id string) error
}

type FilmRepository interface {
	ListAll() ([]entities.Film, error)
	Find(id string) (*entities.Film, error)
	Insert(film *entities.Film) error
	Save(film *entities.Film) error
	Delete(id string) error
	FindGendersById(ids []uint) ([]entities.Gender, error)
	FindDirectorById(id uint) (*entities.Director, error)
}

type SessionRepository interface {
	ListAll() ([]entities.Session, error)
	Find(id string) (*entities.Session, error)
	Insert(session *entities.Session) error
	Save(session *entities.Session) error
	Delete(id string) error
	FindFilmById(id uint) (*entities.Film, error)
	FindByRoomAndTime(room uint, startTime time.Time, endTime time.Time) bool
	FindFirstBeforeTime(room uint, startTime time.Time) *entities.Session
}

type BuyerRepository interface {
	ListAll() ([]entities.Buyer, error)
	Find(id string) (*entities.Buyer, error)
	Insert(buyer *entities.Buyer) error
	Save(buyer *entities.Buyer) error
	Delete(id string) error
}

type TicketRepository interface {
	ListAll() ([]entities.Ticket, error)
	Find(id string) (*entities.Ticket, error)
	Insert(ticket *entities.Ticket) error
	Save(ticket *entities.Ticket) error
	Delete(id string) error
	FindBuyerById(id uint) (*entities.Buyer, error)
	FindSessionById(id uint) (*entities.Session, error)
	FindBySessionIdAndSeat(sessionId uint, seat string) bool
}
