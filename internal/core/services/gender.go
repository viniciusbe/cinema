package services

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.GenderRepository
}

func NewGenderService(r ports.GenderRepository) *Service {
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

func (s *Service) Update(gender *entities.Gender) {
	s.repository.Update(gender)
}

func (s *Service) Create(gender *entities.Gender) {
	s.repository.Insert(gender)
}

func (s *Service) Delete(id string) {
	s.repository.Delete(id)
}
