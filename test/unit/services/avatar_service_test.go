package services

import (
	"1337b04rd/internal/app/common/logger"
	"1337b04rd/internal/app/services"
	"1337b04rd/internal/domain/avatar"
	"context"
	"errors"
	"testing"
	"time"
)

type failingAvatarService struct{}

func (f *failingAvatarService) GetRandomAvatar() (*avatar.Avatar, error) {
	return nil, errors.New("failed to fetch avatar")
}

func TestCreateNewSession_AvatarServiceError(t *testing.T) {
	logger.Init("development")

	repo := &mockSessionRepo{}
	badAvatarSvc := &failingAvatarService{}
	ttl := time.Hour * 24 * 7
	svc := services.NewSessionService(repo, badAvatarSvc, ttl)

	ctx := context.Background()

	sess, err := svc.CreateNew(ctx)
	if err == nil {
		t.Fatal("expected error from CreateNew(), got nil")
	}

	if sess != nil {
		t.Error("expected nil session when avatar service fails")
	}
}
