package service

import "1337B04RD/internal/domain/port"

// internal/domain/service/session_service.go
type SessionService struct {
	sessionRepo   port.SessionRepository
	avatarService port.AvatarService
}

// Методы для работы с сессиями
