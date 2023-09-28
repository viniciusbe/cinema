package buyerrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(gormdb *gorm.DB) ports.BuyerRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Buyer, error) {
	var buyers []entities.Buyer
	err := r.DB.Find(&buyers).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar pagantes -> %w", err)
	}

	return buyers, nil
}

func (r *Repository) Find(id string) (*entities.Buyer, error) {
	var buyer *entities.Buyer
	err := r.DB.First(&buyer, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar pagante -> %w", err)
	}

	return buyer, nil
}

func (r *Repository) Insert(buyer *entities.Buyer) error {
	err := r.DB.Create(&buyer).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir pagante -> %w", err)
	}

	return nil
}

func (r *Repository) Save(buyer *entities.Buyer) error {
	err := r.DB.Save(&buyer).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar pagante -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Buyer{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir pagante -> %w", err)
	}

	return nil
}
