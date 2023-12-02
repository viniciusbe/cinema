package directorrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewNeo4jRepository(gormdb *gorm.DB) ports.DirectorRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Director, error) {
	var directors []entities.Director
	err := r.DB.Find(&directors).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar diretores -> %w", err)
	}

	return directors, nil
}

func (r *Repository) Find(id string) (*entities.Director, error) {
	var director *entities.Director
	err := r.DB.First(&director, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar diretor -> %w", err)
	}

	return director, nil
}

func (r *Repository) Insert(director *entities.Director) error {
	err := r.DB.Create(&director).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir diretor -> %w", err)
	}

	return nil
}

func (r *Repository) Save(director *entities.Director) error {
	err := r.DB.Save(&director).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar diretor -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Director{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir diretor -> %w", err)
	}

	return nil
}
