package filmserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
)

type Service struct {
	repository ports.FilmRepository
}

func New(r ports.FilmRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Film, error) {
	films, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return films, nil
}

func (s *Service) Get(id string) (*entities.Film, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(film *entities.Film) error {
	// if len(gendersID) > 0 {
	// 	genders, err := s.repository.FindGendersById(gendersID)
	// 	film.Genders = genders
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// director, err := s.repository.FindDirectorById(film.DirectorID)
	// if err != nil {
	// 	return err
	// }

	// film.Director = *director
	return s.repository.Save(film)
}

func (s *Service) Create(film *entities.Film) error {
	// genders, genderErr := s.repository.FindGendersById(gendersID)
	// if genderErr != nil {
	// 	return genderErr
	// }
	// film.Genders = genders

	// director, err := s.repository.FindDirectorById(film.DirectorID)
	// if err != nil {
	// 	return err
	// }

	// film.Director = *director
	return s.repository.Insert(film)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) ListOptions() ([]entities.Director, []entities.Gender, error) {
	directors, err := s.repository.GetDirectors()
	if err != nil {
		return nil, nil, err
	}

	genders, err := s.repository.GetGenders()
	if err != nil {
		return nil, nil, err
	}

	return directors, genders, nil
}

type FilmOptions struct {
	directors []entities.Director
	genders   []entities.Gender
}
