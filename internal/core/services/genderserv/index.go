package genderserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.GenderRepository
}

func New(r ports.GenderRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Gender, error) {
	genders, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return genders, nil
}

func (s *Service) Get(id string) (*entities.Gender, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(gender *entities.Gender) error {
	return s.repository.Save(gender)
}

func (s *Service) Create(gender *entities.Gender) error {
	return s.repository.Insert(gender)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
