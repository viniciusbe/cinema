package ports

import "cinema/internal/core/domain/entities"

type GenderRepository interface {
	ListAll() ([]entities.Gender, error)
	Insert(gender *entities.Gender) error
	Update(gender *entities.Gender) error
	Delete(id string)
}
