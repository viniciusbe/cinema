package ports

import "cinema/internal/core/domain"

type GenderRepository interface {
	// Insert(gender *domain.Gender) error
	// Update(gender *domain.Gender) (*domain.Gender, error)
	ListAll() ([]domain.Gender, error)
}
