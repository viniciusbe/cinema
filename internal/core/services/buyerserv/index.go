package buyerserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.BuyerRepository
}

func New(r ports.BuyerRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Buyer, error) {
	buyers, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func (s *Service) Get(id string) (*entities.Buyer, []entities.Ticket, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(buyer *entities.Buyer) error {
	return s.repository.Save(buyer)
}

func (s *Service) Create(buyer *entities.Buyer) error {
	return s.repository.Insert(buyer)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
