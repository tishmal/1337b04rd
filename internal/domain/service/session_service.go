package service

import (
	"1337B04RD/internal/domain/entity"
	"1337B04RD/internal/domain/port"
	"1337B04RD/internal/utils"
	"log/slog"
	"time"
)

type SessionService struct {
	repo   port.SessionRepository
	logger *slog.Logger
}

func NewSessionService(r port.SessionRepository, logger *slog.Logger) *SessionService {
	return &SessionService{
		repo:   r,
		logger: logger}
}

func (s *SessionService) CreateSession() (*entity.Session, error) {
	id := utils.GenerateSessionID()
	session := &entity.Session{
		ID:        id,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	s.logger.Info("Start Create Session",
		"id session", session.ID,
		"user id", session.ExpiresAt,
	)

	err := s.repo.Save(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) GetSession(id string) (*entity.Session, error) {
	return s.repo.Get(id)
}

func (s *SessionService) DeleteSession(id string) error {
	return s.repo.Delete(id)
}
