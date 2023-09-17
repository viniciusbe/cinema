package repositories

import (
	"cinema/internal/core/domain/entities"
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

func (r *Repository) ListAll() ([]entities.Gender, error) {
	var genders []entities.Gender
	err := r.DB.Find(&genders).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar gêneros -> %w", err)
	}

	return genders, nil
}

func (r *Repository) Insert(gender *entities.Gender) error {
	err := r.DB.Create(&gender).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir gênero -> %w", err)
	}

	return nil
}

func (r *Repository) Update(id string, gender *entities.Gender) {

	err := r.DB.Save(&gender).Error

	if err != nil {
		fmt.Println("Genero nao existe")
	}

}

func (r *Repository) Delete(id string) {
	err := r.DB.Delete(&entities.Gender{}, id).Error
	if err != nil {
		fmt.Errorf("Erro do excluir gênero -> %w", err)
	}
}
