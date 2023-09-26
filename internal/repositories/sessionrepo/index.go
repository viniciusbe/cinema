package sessionrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"
	"time"

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
	err := r.DB.Preload("Film").Find(&sessions).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar sessões -> %w", err)
	}

	return sessions, nil
}

func (r *Repository) Find(id string) (*entities.Session, error) {
	var session *entities.Session
	err := r.DB.Preload("Film").First(&session, id).Error
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

func (r *Repository) FindFilmById(id uint) (*entities.Film, error) {
	var film *entities.Film
	err := r.DB.First(&film, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar filme -> %w", err)
	}

	return film, nil
}

func (r *Repository) FindByRoomAndTime(room uint, startTime time.Time, endTime time.Time) bool {
	var session *entities.Session
	rowsAffected := r.DB.Where("room = ? AND time BETWEEN ? AND ?", room, startTime, endTime).First(&session).RowsAffected

	return rowsAffected != 0
}

func (r *Repository) FindFirstBeforeTime(room uint, startTime time.Time) *entities.Session {
	var session *entities.Session
	err := r.DB.Order("time DESC").Where("room = ? AND time < ?", room, startTime).Preload("Film").First(&session).Error
	if err != nil {
		return nil
	}

	return session
}

func (r *Repository) FindByFilmId(id string) ([]entities.Session, error) {
	var sessions []entities.Session
	err := r.DB.Where("film_id = ?", id).Preload("Film").Find(&sessions).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar sessões -> %w", err)
	}

	return sessions, nil
}
