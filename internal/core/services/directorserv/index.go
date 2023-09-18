package directorserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.DirectorRepository
}

func New(r ports.DirectorRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Director, error) {
	directors, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return directors, nil
}

func (s *Service) Get(id string) (*entities.Director, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(director *entities.Director) error {
	return s.repository.Save(director)
}

func (s *Service) Create(director *entities.Director) error {
	return s.repository.Insert(director)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
