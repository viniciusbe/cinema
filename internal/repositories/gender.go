package repositories

import (
	"cinema/internal/core/domain"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormGenderRepository(gormdb *gorm.DB) ports.GenderRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]domain.Gender, error) {
	var genders []domain.Gender
	err := r.DB.Find(&genders).Error
	if err != nil {
		return nil, fmt.Errorf("find -> %w", err)
	}

	return genders, nil

}
