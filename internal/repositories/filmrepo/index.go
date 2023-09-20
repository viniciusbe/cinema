package filmrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"cinema/internal/handlers/cli/gendercli"
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
	err := r.DB.Preload("Genders").Preload("Director").Find(&film, id).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar filme -> %w", err)
	}

	return film, nil
}

func (r *Repository) Insert(film *entities.Film, gendersID []uint) error {
	var genders []entities.Gender
	genderErr := r.DB.Find(&genders, gendersID).Error
	if genderErr != nil {
		return fmt.Errorf("Erro ao buscar genero(s) -> %w", genderErr)
	}

	film.Genders = genders
	err := r.DB.Save(&film).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir filme -> %w", err)
	}

	return nil
}

func (r *Repository) Save(film *entities.Film, gendersID []uint) error {
	var newGenders []entities.Gender
	if len(gendersID) > 0 {
		r.DB.Find(&newGenders, gendersID)
	}
	fmt.Println("---------")
	for _, gender := range film.Genders {
		gendercli.PrintGender(gender)
	}
	fmt.Println("---------")
	for _, gender := range gendersID {
		fmt.Printf("ID %v\n", gender)
	}
	newGenders = append(film.Genders, newGenders[:]...)

	err := r.DB.Save(&film).Error
	r.DB.Model(&film).Association("Genders").Replace(newGenders)
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

// func (r *Repository) findGenders() {

// }
