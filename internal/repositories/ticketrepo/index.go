package ticketrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(gormdb *gorm.DB) ports.TicketRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Ticket, error) {
	var tickets []entities.Ticket
	err := r.DB.Preload("Buyer").Preload("Session").Find(&tickets).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar ingressos -> %w", err)
	}

	return tickets, nil
}

func (r *Repository) Find(id string) (*entities.Ticket, error) {
	var ticket *entities.Ticket
	err := r.DB.Preload("Buyer").Preload("Session").First(&ticket, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar ingresso -> %w", err)
	}

	return ticket, nil
}

func (r *Repository) Insert(ticket *entities.Ticket) error {
	err := r.DB.Save(&ticket).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir ingresso -> %w", err)
	}

	return nil
}

func (r *Repository) Save(ticket *entities.Ticket) error {
	var buyer *entities.Buyer
	buyerErr := r.DB.Find(&buyer, ticket.BuyerID).Error
	if buyerErr != nil {
		return fmt.Errorf("Erro ao buscar pagante -> %w", buyerErr)
	}

	var session *entities.Session
	sessionErr := r.DB.Find(&session, ticket.SessionID).Error
	if sessionErr != nil {
		return fmt.Errorf("Erro ao buscar sessão -> %w", sessionErr)
	}

	ticket.Buyer = *buyer
	ticket.Session = *session
	err := r.DB.Save(&ticket).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar ingresso -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Ticket{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir ingresso -> %w", err)
	}

	return nil
}

func (r *Repository) FindBuyerById(id uint) (*entities.Buyer, error) {
	var buyer *entities.Buyer
	buyerErr := r.DB.First(&buyer, id).Error
	if buyerErr != nil {
		return nil, fmt.Errorf("Erro ao buscar pagante -> %w", buyerErr)
	}

	return buyer, nil
}

func (r *Repository) FindSessionById(id uint) (*entities.Session, error) {
	var session *entities.Session
	sessionErr := r.DB.First(&session, id).Error
	if sessionErr != nil {
		return nil, fmt.Errorf("Erro ao buscar sessão -> %w", sessionErr)
	}

	return session, nil
}

func (r *Repository) FindBySessionIdAndSeat(sessionId uint, seat string) bool {
	rowsAffected := r.DB.Where("session_id = ? AND seat = ?", sessionId, seat).First(&entities.Ticket{}).RowsAffected

	return rowsAffected != 0
}
