package sessionserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"errors"
	"time"
)

type Service struct {
	repository ports.SessionRepository
}

func New(r ports.SessionRepository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListAll() ([]entities.Session, error) {
	sessions, err := s.repository.ListAll()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Service) Get(id string) (*entities.Session, error) {
	return s.repository.Find(id)
}

func (s *Service) Update(session *entities.Session) error {
	film, err := s.repository.FindFilmById(session.FilmID)
	if err != nil {
		return err
	}

	session.Film = *film
	return s.repository.Save(session)
}

func (s *Service) Create(session *entities.Session) error {
	film, err := s.repository.FindFilmById(session.FilmID)
	if err != nil {
		return err
	}
	session.Film = *film

	endTime := session.Time.Add(time.Minute * time.Duration(session.Film.Duration))
	teste := s.repository.FindByRoomAndTime(session.Room, session.Time, endTime)
	if teste {
		return errors.New("Já existe uma sessão nesse horário e sala")
	}
	return s.repository.Insert(session)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
