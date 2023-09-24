package ticketserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"errors"
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
	err := s.Validate(ticket)
	if err != nil {
		return err
	}

	return s.repository.Save(ticket)
}

func (s *Service) Create(ticket *entities.Ticket) error {
	err := s.Validate(ticket)
	if err != nil {
		return err
	}

	return s.repository.Insert(ticket)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) Validate(ticket *entities.Ticket) error {
	buyer, err := s.repository.FindBuyerById(ticket.BuyerID)
	if err != nil {
		return err
	}
	ticket.Buyer = *buyer

	session, err := s.repository.FindSessionById(ticket.SessionID)
	if err != nil {
		return err
	}
	ticket.Session = *session

	isSeatOccupied := s.repository.FindBySessionIdAndSeat(ticket.SessionID, ticket.Seat)
	if isSeatOccupied {
		return errors.New("Assento não disponível.")
	}

	return nil
}
