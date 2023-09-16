package services

import (
	"cinema/internal/core/domain"
	"cinema/internal/core/ports"
	"fmt"
)

type Service struct {
	repository ports.GenderRepository
}

func NewGenderService(repo ports.GenderRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func (serv *Service) ListAll() ([]domain.Gender, error) {
	genders, err := serv.repository.ListAll()
	if err != nil {
		err = fmt.Errorf("create -> %w", err)
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	return genders, nil
}
