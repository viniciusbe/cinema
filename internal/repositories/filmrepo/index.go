package filmrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewGormRepository(gormdb *gorm.DB) ports.FilmRepository {
	return &Repository{
		DB: gormdb,
	}
}

func (r *Repository) ListAll() ([]entities.Film, error) {
	var films []entities.Film
	err := r.DB.Preload("Genders").Preload("Director").Find(&films).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar filmes -> %w", err)
	}

	return films, nil
}

func (r *Repository) Find(id string) (*entities.Film, error) {
	var film *entities.Film
	err := r.DB.Preload("Genders").Preload("Director").First(&film, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar filme -> %w", err)
	}

	return film, nil
}

func (r *Repository) Insert(film *entities.Film) error {
	err := r.DB.Save(&film).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir filme -> %w", err)
	}

	return nil
}

func (r *Repository) Save(film *entities.Film) error {
	r.DB.Model(&film).Association("Genders").Replace(film.Genders)
	err := r.DB.Save(&film).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualizar filme -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	err := r.DB.Delete(&entities.Film{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao excluir filme -> %w", err)
	}

	return nil
}

func (r *Repository) FindGendersById(ids []uint) ([]entities.Gender, error) {
	var genders []entities.Gender
	genderErr := r.DB.Find(&genders, ids).Error
	if genderErr != nil {
		return nil, fmt.Errorf("Erro ao buscar genero(s) -> %w", genderErr)
	}

	return genders, nil
}

func (r *Repository) FindDirectorById(id uint) (*entities.Director, error) {
	var director *entities.Director
	err := r.DB.First(&director, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar diretor -> %w", err)
	}

	return director, nil
}
