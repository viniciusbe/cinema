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
		return errors.New("Erro aqui")
	}
	session.Film = *film

	isTimeClash := s.ValidateTime(session)
	if isTimeClash {
		return errors.New("Já existe uma sessão nesse horário e sala")
	}
	return s.repository.Save(session)
}

func (s *Service) Create(session *entities.Session) error {
	film, err := s.repository.FindFilmById(session.FilmID)
	if err != nil {
		return err
	}
	session.Film = *film

	isTimeClash := s.ValidateTime(session)
	if isTimeClash {
		return errors.New("Erro ao criar sessão -> choque de horários.")
	}
	return s.repository.Insert(session)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) ValidateTime(session *entities.Session) bool {
	isPreviousTimeClash := false

	previousSession := s.repository.FindFirstBeforeTime(session.Room, session.Time)
	if previousSession != nil {
		previousSessionEndTime := previousSession.Time.Add(time.Minute * time.Duration(previousSession.Film.Duration))
		isPreviousTimeClash = previousSessionEndTime.After(session.Time)
	}

	endTime := session.Time.Add(time.Minute * time.Duration(session.Film.Duration))
	isNextTimeClash := s.repository.FindByRoomAndTime(session.Room, session.Time, endTime)

	return isPreviousTimeClash || isNextTimeClash
}

func (s *Service) GetByFilmId(id string) ([]entities.Session, error) {
	sessions, err := s.repository.FindByFilmId(id)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Service) GetByBuyerId(id string) ([]entities.Session, error) {
	sessions, err := s.repository.FindByBuyerId(id)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Service) GetByGenderId(id string) ([]entities.Session, error) {
	sessions, err := s.repository.FindByGenderId(id)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
