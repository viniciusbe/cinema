package genderrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(gormdb *gorm.DB) ports.GenderRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Gender, error) {
	var genders []entities.Gender
	err := r.DB.Find(&genders).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar gêneros -> %w", err)
	}

	return genders, nil
}

func (r *Repository) Find(id string) (*entities.Gender, error) {
	var gender *entities.Gender
	err := r.DB.Find(&gender, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar gênero -> %w", err)
	}

	return gender, nil
}

func (r *Repository) Insert(gender *entities.Gender) error {
	err := r.DB.Create(&gender).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir gênero -> %w", err)
	}

	return nil
}

func (r *Repository) Save(gender *entities.Gender) error {
	err := r.DB.Save(&gender).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar gênero -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Gender{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir gênero -> %w", err)
	}

	return nil
}
