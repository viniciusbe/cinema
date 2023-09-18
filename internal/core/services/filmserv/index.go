package filmserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.FilmRepository
}

func New(r ports.FilmRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Film, error) {
	films, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return films, nil
}

func (s *Service) Get(id string) (*entities.Film, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(film *entities.Film) error {
	return s.repository.Save(film)
}

func (s *Service) Create(film *entities.Film) error {
	return s.repository.Insert(film)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
