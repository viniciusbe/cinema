package sessionrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(gormdb *gorm.DB) ports.SessionRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Session, error) {
	var sessions []entities.Session
	err := r.DB.Find(&sessions).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar sessões -> %w", err)
	}

	return sessions, nil
}

func (r *Repository) Find(id string) (*entities.Session, error) {
	var session *entities.Session
	err := r.DB.Find(&session, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar sessão -> %w", err)
	}

	return session, nil
}

func (r *Repository) Insert(session *entities.Session) error {
	err := r.DB.Create(&session).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir sessão -> %w", err)
	}

	return nil
}

func (r *Repository) Save(session *entities.Session) error {
	err := r.DB.Save(&session).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar sessão -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Session{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir sessão -> %w", err)
	}

	return nil
}
