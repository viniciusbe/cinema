package ports

import "cinema/internal/core/domain/entities"

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
	Insert(film *entities.Film, gendersID []uint) error
	Save(film *entities.Film, gendersID []uint) error
	Delete(id string) error
}

type SessionRepository interface {
	ListAll() ([]entities.Session, error)
	Find(id string) (*entities.Session, error)
	Insert(session *entities.Session) error
	Save(session *entities.Session) error
	Delete(id string) error
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
}
