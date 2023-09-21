package sessionserv

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
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
	return s.repository.Save(session)
}

func (s *Service) Create(session *entities.Session) error {
	return s.repository.Insert(session)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
