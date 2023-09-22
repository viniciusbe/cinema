package ticketserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.TicketRepository
}

func New(r ports.TicketRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Ticket, error) {
	tickets, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *Service) Get(id string) (*entities.Ticket, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(ticket *entities.Ticket) error {
	return s.repository.Save(ticket)
}

func (s *Service) Create(ticket *entities.Ticket) error {
	return s.repository.Insert(ticket)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
